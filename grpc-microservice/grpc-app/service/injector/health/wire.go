//go:build wireinject
// +build wireinject

package health

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	"grpc-microservice/grpc-app/service/health"

	pb "grpc-microservice/grpc-app/protos"

	"github.com/google/wire"
)

func InitializeHealthServerService() (pb.HealthServer, error) {
	wire.Build(
		config.NewPaths,
		config.New,
		config.NewMySQLConnection,
		common.NewLogger,
		health.NewHalthServer,
	)
	return &health.HealthServer{}, nil
}
