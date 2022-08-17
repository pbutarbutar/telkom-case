package interceptor

import (
	"context"
	"grpc-microservice/middleware"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerFunc func(p interface{}) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	Logger     *zap.Logger
	jwtManager middleware.JWTManagerRepo
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jM middleware.JWTManagerRepo, logger *zap.Logger) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager: jM,
		Logger:     logger,
	}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		panicked := true
		defer func() {
			if r := recover(); r != nil || panicked {
				err := recoverFrom(ctx, r, o.recoveryHandlerFunc)
				interceptor.Logger.Error("error", zap.Any("message", err))
			}
		}()

		interceptor.Logger.Sugar().Infow("--> unary interceptor: " + info.FullMethod)
		var userAuthClaim *middleware.User

		if info.FullMethod != "/grpc.health.v1.Health/Check" {
			usr, err := interceptor.authorize(ctx, info.FullMethod)
			if err != nil {
				return nil, err
			}
			userAuthClaim = &usr

			injcContext := NewInjectContext(*userAuthClaim)

			panicked = false
			return injcContext.InjectContextIntercaptorUnary(ctx, req, handler)

		} else {
			panicked = false
			return handler(ctx, req)

		}

	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		panicked := true
		defer func() {
			if r := recover(); r != nil || panicked {
				err := recoverFrom(stream.Context(), r, o.recoveryHandlerFunc)
				interceptor.Logger.Error("error", zap.Any("message", err))
			}
		}()
		interceptor.Logger.Sugar().Infow("--> stream interceptor: " + info.FullMethod)
		var userAuthClaim *middleware.User

		if info.FullMethod != "/grpc.health.v1.Health/Check" {
			usr, err := interceptor.authorize(stream.Context(), info.FullMethod)
			if err != nil {
				return err
			}
			userAuthClaim = &usr

			injcContext := NewInjectContext(*userAuthClaim)

			panicked = false
			return injcContext.InjectContextIntercaptorStream(srv, stream, handler)

		} else {
			panicked = false
			return handler(srv, stream)

		}
	}
}

// Verify verifies the access token string and return a user claim if the token is valid
func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (middleware.User, error) {
	var usr middleware.User
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return usr, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]

	if len(values) < 1 {
		return usr, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	accessToken := strings.TrimPrefix(values[0], "Bearer ")

	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return usr, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if claims.UserID != "" {
		usr.UserID = claims.UserID
		usr.HumanID = claims.HumanID
		usr.TokenAuth = accessToken
		usr.FirstName = claims.FirstName
		usr.FullName = claims.FullName
		usr.Groups = claims.Groups
		usr.Organization = claims.Organization
		usr.IsActive = claims.IsActive
		usr.IsStaff = claims.IsStaff
		return usr, nil
	}

	return usr, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

func recoverFrom(ctx context.Context, p interface{}, r RecoveryHandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
