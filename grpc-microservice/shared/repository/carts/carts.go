package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
)

type CartsRepository interface {
	AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error)
	IsExistProductByCode(ctx context.Context, in *pbCarts.AddProductRequest) (isExist bool, err error)
	UpdateCart(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error)
	DeleteProduct(ctx context.Context, productCode string) (result bool, err error)
	ViewProduct(ctx context.Context) (result []entities.Carts, err error)
}
