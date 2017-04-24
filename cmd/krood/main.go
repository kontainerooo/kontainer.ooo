package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/docker/docker/client"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/containerlifecycle"
	"github.com/kontainerooo/kontainer.ooo/pkg/customercontainer"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiClient "github.com/kontainerooo/kontainer.ooo/pkg/kmi/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

func main() {

	var (
		grpcAddr    = ":8082"
		wsAddr      = ":8081"
		dockerHost  = "http://127.0.0.1:2375"
		isMock      bool
		dbWrapper   abstraction.DB
		dcliWrapper abstraction.DCli
	)

	/* The krood binary can now be given a flag called `--mock`. With this
	 *  option the mock database and mock docker client is used. This is
	 *  in order to simplify testing without a docker daemon and database
	 *  connection. This might later be removed. */
	flag.BoolVar(&isMock, "mock", false, "Determines if a mock DB and docker client should be used.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	if isMock {
		dbWrapper = testutils.NewMockDB()
	} else {
		db, err := gorm.Open("postgres", "host=postgres user=postgres sslmode=disable")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		dbWrapper = abstraction.NewDB(db)
	}

	clientTransport := &http.Client{
		Transport: &http.Transport{},
	}

	if isMock {
		dcliWrapper = testutils.NewMockDCli()
	} else {
		defaultHeaders := map[string]string{}
		dcli, err := client.NewClient(dockerHost, "1.26", clientTransport, defaultHeaders)
		if err != nil {
			panic(err)
		}
		dcliWrapper = abstraction.NewDCLI(dcli)
	}

	var userService user.Service
	userService, err := user.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}
	userService = user.NewTransactionBasedService(userService)

	userEndpoints := makeUserServiceEndpoints(userService)

	var kmiService kmi.Service
	kmiService, err = kmi.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}

	kmiEndpoints := makeKMIServiceEndpoints(kmiService)

	var containerlifecycleService containerlifecycle.Service
	containerlifecycleService = containerlifecycle.NewService(dcliWrapper)

	clsEndpoints := makeCLServiceEndpoints(containerlifecycleService)

	var customercontainerService customercontainer.Service
	customercontainerService = customercontainer.NewService(dcliWrapper)

	ccEndpoint := makeCustomerContainerServiceEndpoints(customercontainerService)

	var routingService routing.Service
	routingService, err = routing.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}

	routingEndpoints := makeRoutingServiceEndpoints(routingService)

	errc := make(chan error)
	ctx := context.Background()

	go startGRPCTransport(ctx, errc, logger, grpcAddr, userEndpoints, kmiEndpoints, clsEndpoints, ccEndpoint, routingEndpoints)

	go startWebsocketTransport(errc, logger, wsAddr, userEndpoints, kmiEndpoints, clsEndpoints, ccEndpoint, routingEndpoints)

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	ke := kmiClient.New(conn, logger)

	customercontainerService.AddKMIClient(ke)
	customercontainerService.AddLogger(logger)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}

func startGRPCTransport(ctx context.Context, errc chan error, logger log.Logger, grpcAddr string, ue user.Endpoints, ke kmi.Endpoints, cle containerlifecycle.Endpoints, cce customercontainer.Endpoints, re routing.Endpoints) {
	logger = log.With(logger, "transport", "gRPC")

	ln, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		errc <- err
		return
	}
	s := grpc.NewServer()

	userServer := user.MakeGRPCServer(ctx, ue, logger)
	pb.RegisterUserServiceServer(s, userServer)

	kmiServer := kmi.MakeGRPCServer(ctx, ke, logger)
	pb.RegisterKMIServiceServer(s, kmiServer)

	clsServer := containerlifecycle.MakeGRPCServer(ctx, cle, logger)
	pb.RegisterContainerLifecycleServiceServer(s, clsServer)

	ccsServer := customercontainer.MakeGRPCServer(ctx, cce, logger)
	pb.RegisterCustomerContainerServiceServer(s, ccsServer)

	routingServer := routing.MakeGRPCServer(ctx, re, logger)
	pb.RegisterRoutingServiceServer(s, routingServer)

	logger.Log("addr", grpcAddr)
	errc <- s.Serve(ln)
}

func startWebsocketTransport(errc chan error, logger log.Logger, wsAddr string, ue user.Endpoints, ke kmi.Endpoints, cle containerlifecycle.Endpoints, cce customercontainer.Endpoints, re routing.Endpoints) {
	logger = log.With(logger, "transport", "ws")
	s := ws.NewServer(ws.ProtocolMap{
		"v1": ws.BasicHandler{},
	}, logger, websocket.Upgrader{
		EnableCompression: true,
	}, ws.NewTokenAuth(), ws.SSLConfig{})

	userService := user.MakeWebsocketService(ue)
	s.RegisterService(userService)

	kmiService := kmi.MakeWebsocketService(ke)
	s.RegisterService(kmiService)

	clsServer := containerlifecycle.MakeWebsocketService(cle)
	s.RegisterService(clsServer)

	ccsServer := customercontainer.MakeWebsocketService(cce)
	s.RegisterService(ccsServer)

	routingServer := routing.MakeWebsocketService(re)
	s.RegisterService(routingServer)

	logger.Log("addr", wsAddr)
	errc <- s.Serve(wsAddr)
}

func makeUserServiceEndpoints(s user.Service) user.Endpoints {
	var createUserEndpoint endpoint.Endpoint
	{
		createUserEndpoint = user.MakeCreateUserEndpoint(s)
	}

	var editUserEndpoint endpoint.Endpoint
	{
		editUserEndpoint = user.MakeEditUserEndpoint(s)
	}

	var changeUsernaemEndpoint endpoint.Endpoint
	{
		changeUsernaemEndpoint = user.MakeChangeUsernameEndpoint(s)
	}

	var deleteUserEndpoint endpoint.Endpoint
	{
		deleteUserEndpoint = user.MakeDeleteUserEndpoint(s)
	}

	var resetPasswordEndpoint endpoint.Endpoint
	{
		resetPasswordEndpoint = user.MakeResetPasswordEndpoint(s)
	}

	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = user.MakeGetUserEndpoint(s)
	}

	return user.Endpoints{
		CreateUserEndpoint:     createUserEndpoint,
		EditUserEndpoint:       editUserEndpoint,
		ChangeUsernameEndpoint: changeUsernaemEndpoint,
		DeleteUserEndpoint:     deleteUserEndpoint,
		ResetPasswordEndpoint:  resetPasswordEndpoint,
		GetUserEndpoint:        getUserEndpoint,
	}
}

func makeKMIServiceEndpoints(s kmi.Service) kmi.Endpoints {
	var AddKMIEndpoint endpoint.Endpoint
	{
		AddKMIEndpoint = kmi.MakeAddKMIEndpoint(s)
	}

	var RemoveKMIEndpoint endpoint.Endpoint
	{
		RemoveKMIEndpoint = kmi.MakeRemoveKMIEndpoint(s)
	}

	var GetKMIEndpoint endpoint.Endpoint
	{
		GetKMIEndpoint = kmi.MakeGetKMIEndpoint(s)
	}

	var KMIEndpoint endpoint.Endpoint
	{
		KMIEndpoint = kmi.MakeKMIEndpoint(s)
	}

	return kmi.Endpoints{
		AddKMIEndpoint:    AddKMIEndpoint,
		RemoveKMIEndpoint: RemoveKMIEndpoint,
		GetKMIEndpoint:    GetKMIEndpoint,
		KMIEndpoint:       KMIEndpoint,
	}
}

func makeCLServiceEndpoints(s containerlifecycle.Service) containerlifecycle.Endpoints {
	var StartContainerEndpoint endpoint.Endpoint
	{
		StartContainerEndpoint = containerlifecycle.MakeStartContainerEndpoint(s)
	}

	var StartCommandEndpoint endpoint.Endpoint
	{
		StartCommandEndpoint = containerlifecycle.MakeStartCommandEndpoint(s)
	}

	var StopContainerEndpoint endpoint.Endpoint
	{
		StopContainerEndpoint = containerlifecycle.MakeStopContainerEndpoint(s)
	}

	return containerlifecycle.Endpoints{
		StartContainerEndpoint: StartContainerEndpoint,
		StartCommandEndpoint:   StartCommandEndpoint,
		StopContainerEndpoint:  StopContainerEndpoint,
	}
}

func makeCustomerContainerServiceEndpoints(s customercontainer.Service) customercontainer.Endpoints {
	var CreateContainerEndpoint endpoint.Endpoint
	{
		CreateContainerEndpoint = customercontainer.MakeCreateContainerEndpoint(s)
	}

	var EditContainerEndpoint endpoint.Endpoint
	{
		EditContainerEndpoint = customercontainer.MakeEditContainerEndpoint(s)
	}

	var RemoveContainerEndpoint endpoint.Endpoint
	{
		RemoveContainerEndpoint = customercontainer.MakeRemoveContainerEndpoint(s)
	}

	var InstancesEndpoint endpoint.Endpoint
	{
		InstancesEndpoint = customercontainer.MakeInstancesEndpoint(s)

	}

	var CreateDockerImageEndpoint endpoint.Endpoint
	{
		CreateDockerImageEndpoint = customercontainer.MakeCreateDockerImageEndpoint(s)

	}

	return customercontainer.Endpoints{
		CreateContainerEndpoint:   CreateContainerEndpoint,
		EditContainerEndpoint:     EditContainerEndpoint,
		RemoveContainerEndpoint:   RemoveContainerEndpoint,
		InstancesEndpoint:         InstancesEndpoint,
		CreateDockerImageEndpoint: CreateDockerImageEndpoint,
	}
}

func makeRoutingServiceEndpoints(s routing.Service) routing.Endpoints {
	var CreateConfigEndpoint endpoint.Endpoint
	{
		CreateConfigEndpoint = routing.MakeCreateConfigEndpoint(s)
	}

	var EditConfigEndpoint endpoint.Endpoint
	{
		EditConfigEndpoint = routing.MakeEditConfigEndpoint(s)
	}

	var GetConfigEndpoint endpoint.Endpoint
	{
		GetConfigEndpoint = routing.MakeGetConfigEndpoint(s)
	}

	var RemoveConfigEndpoint endpoint.Endpoint
	{
		RemoveConfigEndpoint = routing.MakeRemoveConfigEndpoint(s)
	}

	var AddLocationEndpoint endpoint.Endpoint
	{
		AddLocationEndpoint = routing.MakeAddLocationEndpoint(s)
	}

	var RemoveLocationEndpoint endpoint.Endpoint
	{
		RemoveLocationEndpoint = routing.MakeRemoveLocationEndpoint(s)
	}

	var ChangeListenStatementEndpoint endpoint.Endpoint
	{
		ChangeListenStatementEndpoint = routing.MakeChangeListenStatementEndpoint(s)
	}

	var AddServerNameEndpoint endpoint.Endpoint
	{
		AddServerNameEndpoint = routing.MakeAddServerNameEndpoint(s)
	}

	var RemoveServerNameEndpoint endpoint.Endpoint
	{
		RemoveServerNameEndpoint = routing.MakeRemoveServerNameEndpoint(s)
	}

	var ConfigurationsEndpoint endpoint.Endpoint
	{
		ConfigurationsEndpoint = routing.MakeConfigurationsEndpoint(s)
	}

	return routing.Endpoints{
		CreateConfigEndpoint:          CreateConfigEndpoint,
		EditConfigEndpoint:            EditConfigEndpoint,
		GetConfigEndpoint:             GetConfigEndpoint,
		RemoveConfigEndpoint:          RemoveConfigEndpoint,
		AddLocationEndpoint:           AddLocationEndpoint,
		RemoveLocationEndpoint:        RemoveLocationEndpoint,
		ChangeListenStatementEndpoint: ChangeListenStatementEndpoint,
		AddServerNameEndpoint:         AddServerNameEndpoint,
		RemoveServerNameEndpoint:      RemoveServerNameEndpoint,
		ConfigurationsEndpoint:        ConfigurationsEndpoint,
	}
}
