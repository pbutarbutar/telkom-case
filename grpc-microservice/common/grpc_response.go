package common

import (
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
)

func MapCartsProductResponse(cart entities.Carts) *pbCarts.ProductResponse {
	res := pbCarts.ProductResponse{}
	res.ProductCode = cart.ProductCode
	res.ProductName = cart.ProductName
	res.Quantity = cart.Quantity
	return &res
}

func MapCartsProductsResponse(cart []entities.Carts) (response *pbCarts.ProductsListResponse) {
	resProd := pbCarts.ProductsListResponse{}
	for _, res := range cart {
		res := pbCarts.ProductResponse{
			ProductCode: res.ProductCode,
			ProductName: res.ProductName,
			Quantity:    res.Quantity,
		}
		resProd.ProductsList = append(resProd.ProductsList, &res)
	}

	return &resProd
}
