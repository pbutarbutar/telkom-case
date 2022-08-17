package main

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	pbHealth "grpc-microservice/grpc-app/protos"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	pbUsers "grpc-microservice/grpc-app/protos/users"
	cartsInjector "grpc-microservice/grpc-app/service/injector/carts"
	healthInjector "grpc-microservice/grpc-app/service/injector/health"
	usersInjector "grpc-microservice/grpc-app/service/injector/users"
	"grpc-microservice/grpc-app/service/interceptor"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	mwInjector "grpc-microservice/middleware/injector"
)

var cfg config.Config

func init() {
	v := config.NewPaths()
	cfg = config.New(v...)
}

func main() {
	logger := common.NewLogger()
	jwtManager, _ := mwInjector.InitializeJWTMiddleware()
	interceptor := interceptor.NewAuthInterceptor(jwtManager, logger)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	listener, err := net.Listen("tcp", cfg.Get("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(grpcServer)

	healthService, _ := healthInjector.InitializeHealthServerService()
	usersService, _ := usersInjector.InitializeUsersService()
	cartService, _ := cartsInjector.InitializeCartsService()

	pbHealth.RegisterHealthServer(grpcServer, healthService)
	pbUsers.RegisterUserServiceServer(grpcServer, usersService)
	pbCarts.RegisterCartsServiceServer(grpcServer, cartService)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	logger.Info("GRPC service started on port " + cfg.Get("GRPC_PORT"))

	<-done

	logger.Info("service is going to stop")
	grpcServer.Stop()
	logger.Info("service exited properly")
}
