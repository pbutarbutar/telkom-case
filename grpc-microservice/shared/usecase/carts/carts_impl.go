package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
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

func (c CartsUseCaseImpl) AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error) {
	Isxist, carts, err := c.Repo.IsExistProductByCode(ctx, in)
	if err != nil {
		c.Logger.Error("CartsUseCaseImpl#AddProduct#IsExistProductByCode error: ", zap.Any("message", err))
		return
	}

	if Isxist {
		in.Quantity = in.Quantity + carts.Quantity
		result, err = c.Repo.UpdateCart(ctx, in)
		if err != nil {
			c.Logger.Error("CartsUseCaseImpl#AddProduct#UpdateCart error: ", zap.Any("message", err))
			return
		}
	} else {
		result, err = c.Repo.AddProduct(ctx, in)
		if err != nil {
			c.Logger.Error("CartsUseCaseImpl#AddProduct#AddProduct error: ", zap.Any("message", err))
			return
		}
	}
	return
}

func (c CartsUseCaseImpl) DeleteProduct(ctx context.Context, productCode string) (result bool, err error) {
	result, err = c.Repo.DeleteProduct(ctx, productCode)
	if err != nil {
		c.Logger.Error("CartsUseCaseImpl#DeleteProduct#DeleteProduct error: ", zap.Any("message", err))
		return
	}
	return
}

func (c CartsUseCaseImpl) ViewProduct(ctx context.Context) (result []entities.Carts, err error) {
	result, err = c.Repo.ViewProducts(ctx)
	if err != nil {
		c.Logger.Error("CartsUseCaseImpl#ViewProduct#ViewProducts error: ", zap.Any("message", err))
		return
	}
	return
}
