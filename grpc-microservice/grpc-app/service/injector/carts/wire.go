//go:build wireinject
// +build wireinject

package injector

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	repoCarts "grpc-microservice/shared/repository/carts"
	uCarts "grpc-microservice/shared/usecase/carts"
	"grpc-microservice/shared/utils"

	"github.com/google/wire"

	pbCarts "grpc-microservice/grpc-app/protos/carts"
	serviceCarts "grpc-microservice/grpc-app/service/carts"
)

func InitializeCartsService() (pbCarts.CartsServiceServer, error) {
	wire.Build(
		config.NewPaths,
		config.New,
		config.NewMySQLConnection,
		common.NewLogger,
		utils.NewUtility,
		serviceCarts.NewCartsService,
		repoCarts.NewCartsRepositoryImpl,
		uCarts.NewCartsUseCaseImpl,
	)
	return &serviceCarts.CartsService{}, nil
}
