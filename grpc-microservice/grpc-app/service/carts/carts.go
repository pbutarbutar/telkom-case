package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	ucsCarts "grpc-microservice/shared/usecase/carts"

	"go.uber.org/zap"
)

type CartsService struct {
	pbCarts.UnimplementedCartsServiceServer
	CartUseCase ucsCarts.CartsUseCase
	Logger      *zap.Logger
}

func NewCartsService(uc ucsCarts.CartsUseCase, logger *zap.Logger) pbCarts.CartsServiceServer {
	return &CartsService{
		CartUseCase: uc,
		Logger:      logger,
	}
}

func (u CartsService) AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (*pbCarts.ProductResponse, error) {
	return nil, nil
}

func (u CartsService) DeleteProduct(ctx context.Context, in *pbCarts.DeleteProductRequest) (*pbCarts.ActionResponse, error) {
	return nil, nil
}

func (u CartsService) ViewProduct(ctx context.Context, in *pbCarts.Empty) (*pbCarts.ProductsListResponse, error) {
	return nil, nil
}
