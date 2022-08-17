//go:build wireinject
// +build wireinject

package injector

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	"grpc-microservice/middleware"

	"github.com/google/wire"
)

func InitializeJWTMiddleware() (*middleware.JWTManager, error) {
	wire.Build(
		config.NewPaths,
		config.New,
		config.NewJwtConfig,
		common.NewLogger,
		middleware.NewJWTManager,
	)
	return &middleware.JWTManager{}, nil
}
