// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package health

import (
	"grpc-microservice/common"
	"grpc-microservice/config"
	"grpc-microservice/grpc-app/protos"
	"grpc-microservice/grpc-app/service/health"
)

// Injectors from wire.go:

func InitializeHealthServerService() (__.HealthServer, error) {
	v := config.NewPaths()
	configConfig := config.New(v...)
	db := config.NewMySQLConnection(configConfig)
	logger := common.NewLogger()
	healthServer := health.NewHalthServer(db, logger)
	return healthServer, nil
}
