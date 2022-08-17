package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
)

type CartsRepository interface {
	AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error)
	IsExistProductByCode(ctx context.Context, in *pbCarts.AddProductRequest) (isExist bool, carts *entities.Carts, err error)
	UpdateCart(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error)
	DeleteProduct(ctx context.Context, productCode string) (result bool, err error)
	ViewProducts(ctx context.Context) (result []entities.Carts, err error)
}
