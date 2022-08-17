package users

import (
	"context"
	"grpc-microservice/shared/models/entities"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type UsersRepositoryImpl struct {
	DB     *bun.DB
	Logger *zap.Logger
}

func NewUsersRepositoryImpl(db *bun.DB, logger *zap.Logger) UsersRepository {
	return &UsersRepositoryImpl{
		DB:     db,
		Logger: logger,
	}
}

func (r *UsersRepositoryImpl) CreateUser(ctx context.Context, u entities.Users) (e entities.Users) {
	return
}

func (r *UsersRepositoryImpl) Login(ctx context.Context) (e entities.Users) {
	return
}
