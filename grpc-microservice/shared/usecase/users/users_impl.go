package users

import (
	"context"
	"grpc-microservice/shared/models/entities"
	repo_user "grpc-microservice/shared/repository/users"
	"grpc-microservice/shared/utils"

	"go.uber.org/zap"
)

type UsersUseCaseImpl struct {
	Repo   repo_user.UsersRepository
	Utils  utils.Utility
	Logger *zap.Logger
}

func NewUsersUseCaseImpl(repo repo_user.UsersRepository, utils utils.Utility, logger *zap.Logger) UsersUseCase {
	return &UsersUseCaseImpl{
		Repo:   repo,
		Utils:  utils,
		Logger: logger,
	}
}

func (r *UsersUseCaseImpl) CreateUser(ctx context.Context, u entities.Users) (e entities.Users) {
	return
}

func (r *UsersUseCaseImpl) Login(ctx context.Context) (e entities.Users) {
	return
}
