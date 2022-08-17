package config

import (
	"grpc-microservice/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.HTTPErrorHandler,
	}
}
