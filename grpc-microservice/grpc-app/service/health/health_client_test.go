package health

import (
	"context"
	"database/sql"
	"grpc-microservice/common"
	"grpc-microservice/config"
	pb "grpc-microservice/grpc-app/protos"
	"grpc-microservice/grpc-app/service/interceptor"
	"grpc-microservice/middleware"
	"net"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var header metadata.MD
var ctxOg context.Context

var _ = Describe("Test Client GRGP - Health Services", func() {

	BeforeEach(func() {
		header = metadata.New(map[string]string{"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2Nzg5MCIsImh1bWFuX2lkIjoiSm9obiBEb2UiLCJJc0FjdGl2ZSI6dHJ1ZX0.j3ojJLhR1r2nJB2WooDB9L4gOMGEx1f3k4vDRC-Z0YQ"})
		ctxOg = metadata.NewOutgoingContext(context.Background(), header)
	})

	It("Test Client - Health Access PING it's work", func() {
		sqldb, _ := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
		sqldb.SetMaxOpenConns(1)
		DB = bun.NewDB(sqldb, sqlitedialect.New())
		DeferCleanup(DB.Close)
		serverAddress := startTestOrdersServer()
		healthClient := newTestOrdersClient(serverAddress)
		result, err := healthClient.Check(ctxOg, &pb.HealthCheckRequest{Service: "Test"})
		Expect(err).To(BeNil())

		strResult := result.GetStatus()
		expectedVal := pb.HealthCheckResponse_ServingStatus(1)
		Expect(strResult).To(Equal(expectedVal))

	})

	It("Test Client - Health Access PING does'nt work", func() {
		sqldb, _ := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
		sqldb.SetMaxOpenConns(1)
		DB = bun.NewDB(sqldb, sqlitedialect.New())
		DB.Close() //make closed

		serverAddress := startTestOrdersServer()
		healthClient := newTestOrdersClient(serverAddress)
		result, err := healthClient.Check(ctxOg, &pb.HealthCheckRequest{Service: "Test"})
		Expect(err).NotTo(BeNil())
		Expect(result).To(BeNil())
	})

})

func startTestOrdersServer() string {
	cJW := config.ConfigJWTEnv{
		JwtSecret:        "hkajhdjdkjdjadje249482kjdadadda",
		DurationTRefresh: 15 * time.Minute,
	}
	logger := common.NewLogger()
	jwtManager := middleware.NewJWTManager(&cJW, logger)
	interceptor := interceptor.NewAuthInterceptor(jwtManager, logger)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	newOrders := NewHalthServer(DB, logger)

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterHealthServer(grpcServer, newOrders)

	listener, err := net.Listen("tcp", ":0") // random available port
	Expect(err).To(BeNil())

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestOrdersClient(serverAddress string) pb.HealthClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	Expect(err).To(BeNil())
	return pb.NewHealthClient(conn)
}
