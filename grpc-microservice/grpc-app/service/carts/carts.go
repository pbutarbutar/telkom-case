package carts

import (
	"context"
	"fmt"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	ucsCarts "grpc-microservice/shared/usecase/carts"

	"grpc-microservice/common"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	fmt.Println("NYAMPE")

	result, err := u.CartUseCase.AddProduct(ctx, in)

	fmt.Println("USECASE")

	if err != nil {
		u.Logger.Error("CartsService[AddProduct] error", zap.Any("message", err))
		return nil, status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	resp := common.MapCartsProductResponse(result)
	return resp, nil
}

func (u CartsService) DeleteProduct(ctx context.Context, in *pbCarts.DeleteProductRequest) (*pbCarts.ActionResponse, error) {
	result, err := u.CartUseCase.DeleteProduct(ctx, in.GetProductCode())

	if err != nil {
		u.Logger.Error("CartsService[DeleteProduct] error", zap.Any("message", err))
		return nil, status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	if result {
		return &pbCarts.ActionResponse{
			Success: result,
			Msg:     "Successfully",
		}, nil
	} else {
		return &pbCarts.ActionResponse{
			Success: result,
			Msg:     "Not Successfully",
		}, nil
	}

}

func (u CartsService) ViewProduct(ctx context.Context, in *pbCarts.Empty) (*pbCarts.ProductsListResponse, error) {
	result, err := u.CartUseCase.ViewProduct(ctx)

	if err != nil {
		u.Logger.Error("CartsService[ViewProduct] error", zap.Any("message", err))
		return nil, status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	resp := common.MapCartsProductsResponse(result)

	return resp, nil
}
