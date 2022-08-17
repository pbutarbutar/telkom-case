// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	"grpc-microservice/grpc-app/protos/users"
	users3 "grpc-microservice/grpc-app/service/users"
	"grpc-microservice/shared/repository/users"
	users2 "grpc-microservice/shared/usecase/users"
	"grpc-microservice/shared/utils"
)

// Injectors from wire.go:

func InitializeUsersService() (__.UserServiceServer, error) {
	v := config.NewPaths()
	configConfig := config.New(v...)
	db := config.NewMySQLConnection(configConfig)
	logger := common.NewLogger()
	usersRepository := users.NewUsersRepositoryImpl(db, logger)
	utility := utils.NewUtility()
	usersUseCase := users2.NewUsersUseCaseImpl(usersRepository, utility, logger)
	userServiceServer := users3.NewOrderService(usersUseCase, logger)
	return userServiceServer, nil
}