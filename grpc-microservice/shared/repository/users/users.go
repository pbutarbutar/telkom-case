package users

import (
	"context"
	"grpc-microservice/shared/models/entities"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, u entities.Users) entities.Users
	Login(ctx context.Context) entities.Users
}
