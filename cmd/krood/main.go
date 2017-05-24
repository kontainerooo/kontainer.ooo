package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/opencontainers/runc/libcontainer"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	"github.com/kontainerooo/kontainer.ooo/pkg/kentheguru"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	"github.com/kontainerooo/kontainer.ooo/pkg/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
	_ "github.com/opencontainers/runc/libcontainer/nsenter"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		runtime.GOMAXPROCS(1)
		runtime.LockOSThread()
		factory, err := libcontainer.New("")
		if err != nil {
			panic(err)
		}

		if err := factory.StartInitialization(); err != nil {
			panic(err)
		}
		panic("--this line should have never been executed, congratulations--")
	}
}

func main() {

	var (
		grpcAddr     = ":8082"
		wsAddr       = ":8083"
		wsAddrSecure = ":8084"
		bcryptCost   = 15
		isMock       bool
		dbWrapper    abstraction.DB
	)

	/* The krood binary can now be given a flag called `--mock`. With this
	 *  option the mock database is used. This is
	 *  in order to simplify testing without a database
	 *  connection. This might later be removed. */
	flag.BoolVar(&isMock, "mock", false, "Determines if a mock DB should be used.")
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
		db, err := gorm.Open("postgres", "host=postgres database=postgres user=kroo password=kroo sslmode=disable")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		dbWrapper = abstraction.NewDB(db)
	}

	var userService user.Service
	userService, err := user.NewService(dbWrapper, bcryptCost)
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

	var routingService routing.Service
	routingService, err = routing.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}

	routingEndpoints := makeRoutingServiceEndpoints(routingService)

	factory, err := libcontainer.New("/var/lib/kontainerooo/container", libcontainer.Cgroupfs, libcontainer.InitArgs(os.Args[0], "init"))
	if err != nil {
		panic(err)
		return
	}

	var containerService container.Service
	containerService, err = container.NewService(factory, dbWrapper, &kmiEndpoints, logger)
	if err != nil {
		panic(err)
	}

	containerServiceEndpoints := makeContainerServiceEndpoints(containerService)

	errc := make(chan error)
	ctx := context.Background()

	go startGRPCTransport(ctx, errc, logger, grpcAddr, userEndpoints, kmiEndpoints, routingEndpoints, containerServiceEndpoints)

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	kenTheGuruService := kentheguru.NewService(
		"bu", "bu", "bu", // TODO: generate keys
		websocket.Upgrader{
			EnableCompression: true,
		},
		ws.SSLConfig{
			Addr: wsAddrSecure,
			// TODO: generate certificate and key
		},
		userEndpoints, kmiEndpoints, containerServiceEndpoints, routingEndpoints,
	)

	go kenTheGuruService.StartWebsocketTransport(errc, logger, wsAddr)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}

func startGRPCTransport(ctx context.Context, errc chan error, logger log.Logger, grpcAddr string, ue user.Endpoints, ke kmi.Endpoints, re routing.Endpoints, ce container.Endpoints) {
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

	routingServer := routing.MakeGRPCServer(ctx, re, logger)
	pb.RegisterRoutingServiceServer(s, routingServer)

	containerServer := container.MakeGRPCServer(ctx, ce, logger)
	pb.RegisterContainerServiceServer(s, containerServer)

	logger.Log("addr", grpcAddr)
	errc <- s.Serve(ln)
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

	var checkLoginCredentialsEndpoint endpoint.Endpoint
	{
		checkLoginCredentialsEndpoint = user.MakeCheckLoginCredentialsEndpoint(s)
	}

	return user.Endpoints{
		CreateUserEndpoint:            createUserEndpoint,
		EditUserEndpoint:              editUserEndpoint,
		ChangeUsernameEndpoint:        changeUsernaemEndpoint,
		DeleteUserEndpoint:            deleteUserEndpoint,
		ResetPasswordEndpoint:         resetPasswordEndpoint,
		GetUserEndpoint:               getUserEndpoint,
		CheckLoginCredentialsEndpoint: checkLoginCredentialsEndpoint,
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

func makeContainerServiceEndpoints(s container.Service) container.Endpoints {

	var CreateContainerEndpoint endpoint.Endpoint
	{
		CreateContainerEndpoint = container.MakeCreateContainerEndpoint(s)
	}
	var RemoveContainerEndpoint endpoint.Endpoint
	{
		RemoveContainerEndpoint = container.MakeRemoveContainerEndpoint(s)
	}
	var InstancesEndpoint endpoint.Endpoint
	{
		InstancesEndpoint = container.MakeInstancesEndpoint(s)
	}
	var StopContainerEndpoint endpoint.Endpoint
	{
		StopContainerEndpoint = container.MakeStopContainerEndpoint(s)
	}
	var IsRunningEndpoint endpoint.Endpoint
	{
		IsRunningEndpoint = container.MakeIsRunningEndpoint(s)
	}
	var ExecuteEndpoint endpoint.Endpoint
	{
		ExecuteEndpoint = container.MakeExecuteEndpoint(s)
	}

	return container.Endpoints{
		CreateContainerEndpoint: CreateContainerEndpoint,
		RemoveContainerEndpoint: RemoveContainerEndpoint,
		InstancesEndpoint:       InstancesEndpoint,
		StopContainerEndpoint:   StopContainerEndpoint,
		IsRunningEndpoint:       IsRunningEndpoint,
		ExecuteEndpoint:         ExecuteEndpoint,
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
