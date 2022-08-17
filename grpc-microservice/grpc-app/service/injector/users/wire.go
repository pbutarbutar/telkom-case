//go:build wireinject
// +build wireinject

package injector

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	repoUsers "grpc-microservice/shared/repository/users"
	ucsUsers "grpc-microservice/shared/usecase/users"
	"grpc-microservice/shared/utils"

	"github.com/google/wire"

	pb "grpc-microservice/grpc-app/protos/users"
	serviceUsers "grpc-microservice/grpc-app/service/users"
)

func InitializeUsersService() (pb.UserServiceServer, error) {
	wire.Build(
		config.NewPaths,
		config.New,
		config.NewMySQLConnection,
		common.NewLogger,
		utils.NewUtility,
		serviceUsers.NewOrderService,
		repoUsers.NewUsersRepositoryImpl,
		ucsUsers.NewUsersUseCaseImpl,
	)
	return &serviceUsers.UsersService{}, nil
}
