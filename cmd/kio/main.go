package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ttdennis/kontainer.io/pkg/abstraction"
	"github.com/ttdennis/kontainer.io/pkg/pb"
	"github.com/ttdennis/kontainer.io/pkg/user"
)

func main() {

	var (
		grpcAddr = ":8082"
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	db, err := gorm.Open("postgres", "host=postgres user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dbWrapper := abstraction.NewDB(db)

	var userService user.Service
	{
		userService, err = user.NewService(dbWrapper)
		if err != nil {
			panic(err)
		}
		userService = user.NewTransactionBasedService(userService)
	}

	userEndpoints := makeUserServiceEndpoints(userService)

	errc := make(chan error)
	ctx := context.Background()

	go startGRPCTransport(ctx, errc, logger, grpcAddr, userEndpoints)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}

func startGRPCTransport(ctx context.Context, errc chan error, logger log.Logger, grpcAddr string, ue user.Endpoints) {
	logger = log.NewContext(logger).With("transport", "gRPC")

	ln, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		errc <- err
		return
	}

	srv := user.MakeGRPCServer(ctx, ue, logger)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, srv)

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

	return user.Endpoints{
		CreateUserEndpoint:     createUserEndpoint,
		EditUserEndpoint:       editUserEndpoint,
		ChangeUsernameEndpoint: changeUsernaemEndpoint,
		DeleteUserEndpoint:     deleteUserEndpoint,
		ResetPasswordEndpoint:  resetPasswordEndpoint,
		GetUserEndpoint:        getUserEndpoint,
	}
}
