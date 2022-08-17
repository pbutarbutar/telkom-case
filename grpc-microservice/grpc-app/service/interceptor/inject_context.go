package interceptor

import (
	"context"
	"fmt"
	"grpc-microservice/middleware"

	"google.golang.org/grpc"
)

type UserAuthCtx struct {
	UserAuthenticated middleware.User
}

type secretType string

const SecretKey = secretType("ctx-key-grpc")

func NewInjectContext(userAuthClaim middleware.User) *UserAuthCtx {
	return &UserAuthCtx{
		UserAuthenticated: userAuthClaim,
	}

}

func (u UserAuthCtx) InjectContextIntercaptorUnary(ctx context.Context, req interface{}, handler grpc.UnaryHandler) (interface{}, error) {
	newCtx := context.WithValue(ctx, SecretKey, u)
	return handler(newCtx, req)

}

func (u UserAuthCtx) InjectContextIntercaptorStream(srv interface{}, stream grpc.ServerStream, handler grpc.StreamHandler) error {
	newStream := WrapServerStream(stream)
	newStream.WrappedContext = context.WithValue(stream.Context(), SecretKey, u)
	return handler(srv, newStream)

}

func GetUserAuthCtx(ctx context.Context) (*UserAuthCtx, error) {
	cfg, ok := ctx.Value(SecretKey).(UserAuthCtx)
	if !ok {
		return nil, fmt.Errorf("user auth not found")
	}
	return &cfg, nil

}
