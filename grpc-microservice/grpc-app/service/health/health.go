package health

import (
	"context"
	pb "grpc-microservice/grpc-app/protos"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HealthServer struct {
	pb.UnimplementedHealthServer
	DB     *bun.DB
	Logger *zap.Logger
}

func (h HealthServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if err := h.DB.Ping(); err != nil {
		h.Logger.Error("HealthServer[Check] error", zap.Any("message", err))
		return &pb.HealthCheckResponse{
			Status: 2,
		}, status.Errorf(codes.Internal, "")
	}
	return &pb.HealthCheckResponse{
		Status: 1,
	}, nil
}

func NewHalthServer(db *bun.DB, logger *zap.Logger) pb.HealthServer {
	return &HealthServer{DB: db, Logger: logger}
}
