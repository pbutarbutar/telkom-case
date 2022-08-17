package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
)

type CartsUseCase interface {
	AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error)
	DeleteProduct(ctx context.Context, productCode string) (result bool, err error)
	ViewProduct(ctx context.Context) (result []entities.Carts, err error)
}
