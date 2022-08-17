package exception

import (
	"grpc-microservice/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HTTPErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(common.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(common.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
