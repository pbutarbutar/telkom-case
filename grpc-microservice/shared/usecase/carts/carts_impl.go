package carts

import (
	repo_carts "grpc-microservice/shared/repository/carts"
	"grpc-microservice/shared/utils"

	"go.uber.org/zap"
)

type CartsUseCaseImpl struct {
	Repo   repo_carts.CartsRepository
	Utils  utils.Utility
	Logger *zap.Logger
}

func NewCartsUseCaseImpl(repo repo_carts.CartsRepository, utils utils.Utility, logger *zap.Logger) CartsUseCase {
	return &CartsUseCaseImpl{
		Repo:   repo,
		Utils:  utils,
		Logger: logger,
	}
}
