package users

import (
	"context"
	pb "grpc-microservice/grpc-app/protos/users"
	ucsUsers "grpc-microservice/shared/usecase/users"

	"go.uber.org/zap"
)

type UsersService struct {
	pb.UnimplementedUserServiceServer
	UsersUseCase ucsUsers.UsersUseCase
	Logger       *zap.Logger
}

func NewOrderService(uc ucsUsers.UsersUseCase, logger *zap.Logger) pb.UserServiceServer {
	return &UsersService{
		UsersUseCase: uc,
		Logger:       logger,
	}
}

func (u UsersService) CreateUser(ctx context.Context, req *pb.NewUser) (*pb.User, error) {

	return nil, nil
}
