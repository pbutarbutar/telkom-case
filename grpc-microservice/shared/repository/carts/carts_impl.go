package carts

import (
	"context"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
	"time"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type CartsRepositoryImpl struct {
	DB     *bun.DB
	Logger *zap.Logger
}

func NewCartsRepositoryImpl(db *bun.DB, logger *zap.Logger) CartsRepository {
	return &CartsRepositoryImpl{
		DB:     db,
		Logger: logger,
	}
}

func (c CartsRepositoryImpl) IsExistProductByCode(ctx context.Context, in *pbCarts.AddProductRequest) (isExist bool, carts *entities.Carts, err error) {
	cartsMdl := entities.Carts{
		ProductCode: in.ProductCode,
		ProductName: in.ProductName,
		Quantity:    in.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	isExist, err = c.DB.NewSelect().Model(&cartsMdl).WherePK().Exists(ctx)
	if err != nil {
		c.Logger.Error("CartsRepositoryImpl#IsExistProductByCode error: ", zap.Any("message", err))
		return
	}

	return isExist, &cartsMdl, nil
}

func (c CartsRepositoryImpl) AddProduct(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error) {
	cartsMdl := entities.Carts{
		ProductCode: in.ProductCode,
		ProductName: in.ProductName,
		Quantity:    in.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	res, err := c.DB.NewInsert().Model(&cartsMdl).Returning("*").Exec(ctx)
	if err != nil {
		c.Logger.Error("CartsRepositoryImpl#AddProduct error: ", zap.Any("message", err))
		return
	}

	resEnd, _ := res.RowsAffected()
	if resEnd > 0 {
		result = cartsMdl
	}
	return
}

func (c CartsRepositoryImpl) UpdateCart(ctx context.Context, in *pbCarts.AddProductRequest) (result entities.Carts, err error) {

	cartsMdl := entities.Carts{
		ProductCode: in.ProductCode,
		ProductName: in.ProductName,
		Quantity:    in.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	res, err := c.DB.NewUpdate().Model(&cartsMdl).WherePK().Exec(ctx)
	if err != nil {
		c.Logger.Error("CartsRepositoryImpl#UpdateCart error: ", zap.Any("message", err))
		return
	}

	resEnd, _ := res.RowsAffected()
	if resEnd > 0 {
		result = cartsMdl
	}
	return
}

func (c CartsRepositoryImpl) DeleteProduct(ctx context.Context, productCode string) (result bool, err error) {
	cartsMdl := entities.Carts{
		ProductCode: productCode,
	}
	_, err = c.DB.NewDelete().Model(&cartsMdl).WherePK().Returning("*").Exec(ctx)
	if err != nil {
		c.Logger.Error("CartsRepositoryImpl#DeleteProduct error: ", zap.Any("message", err))
		return
	}

	return true, nil
}

func (c CartsRepositoryImpl) ViewProducts(ctx context.Context) (result []entities.Carts, err error) {
	resultData := []entities.Carts{}
	err = c.DB.NewSelect().Model(&resultData).Scan(ctx)
	if err != nil {
		c.Logger.Error("CartsRepositoryImpl#ViewProduct error: ", zap.Any("message", err))
		return
	}
	return resultData, nil
}
