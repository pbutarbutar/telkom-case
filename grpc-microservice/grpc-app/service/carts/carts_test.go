package carts

import (
	"context"
	"fmt"
	"grpc-microservice/common"
	"grpc-microservice/config"
	"grpc-microservice/middleware"
	"net"
	"testing"
	"time"

	pb "grpc-microservice/grpc-app/protos/carts"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/grpc-app/service/interceptor"
	"grpc-microservice/shared/models/entities"
	mock_cartsUsecase "grpc-microservice/shared/usecase/carts/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var ctxOg context.Context
var ctxServiceToken context.Context
var header metadata.MD

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GRPC Carts Service Suite")
}

var _ = Describe("Checking Carts Grpc Service", Ordered, func() {

	var (
		grpcService pb.CartsServiceServer
		ctrl        *gomock.Controller
		mockCarttU  *mock_cartsUsecase.MockCartsUseCase
	)

	any := gomock.Any()

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockCarttU = mock_cartsUsecase.NewMockCartsUseCase(ctrl)
		header = metadata.New(map[string]string{"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2Nzg5MCIsImh1bWFuX2lkIjoiSm9obiBEb2UiLCJJc0FjdGl2ZSI6dHJ1ZX0.j3ojJLhR1r2nJB2WooDB9L4gOMGEx1f3k4vDRC-Z0YQ"})
		ctxOg = metadata.NewOutgoingContext(context.Background(), header)
		ctxServiceToken = metadata.NewOutgoingContext(context.Background(), header)

	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("AddProduct", func() {
		It("SUCCESS", func() {
			mockCarttU.EXPECT().AddProduct(any, any).Return(
				entities.Carts{
					ProductCode: "P001",
					ProductName: "Product 1",
					Quantity:    4,
				},
				nil,
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			result, err := cartRpcClient.AddProduct(ctxServiceToken, &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    4,
			})

			Expect(err).To(BeNil())
			Expect(result.GetProductCode()).To(Equal("P001"))
			Expect(result.GetProductName()).To(Equal("Product 1"))
			Expect(result.GetQuantity()).To(Equal(int32(4)))

		})

		It("Error", func() {
			mockCarttU.EXPECT().AddProduct(any, any).Return(
				entities.Carts{},
				fmt.Errorf("error"),
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			_, err := cartRpcClient.AddProduct(ctxServiceToken, &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    4,
			})

			Expect(err).NotTo(BeNil())

		})

	})

	Context("DeleteProduct", func() {
		It("SUCCESS Is true", func() {
			mockCarttU.EXPECT().DeleteProduct(any, any).Return(
				true,
				nil,
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			result, err := cartRpcClient.DeleteProduct(ctxServiceToken, &pbCarts.DeleteProductRequest{
				ProductCode: "P001",
			})

			Expect(err).To(BeNil())
			Expect(result.GetSuccess()).To(BeTrue())
		})

		It("SUCCESS Is False", func() {
			mockCarttU.EXPECT().DeleteProduct(any, any).Return(
				false,
				nil,
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			result, err := cartRpcClient.DeleteProduct(ctxServiceToken, &pbCarts.DeleteProductRequest{
				ProductCode: "P001",
			})

			Expect(err).To(BeNil())
			Expect(result.GetSuccess()).To(BeFalse())
		})

		It("Error", func() {
			mockCarttU.EXPECT().DeleteProduct(any, any).Return(
				false,
				fmt.Errorf("error"),
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			_, err := cartRpcClient.DeleteProduct(ctxServiceToken, &pbCarts.DeleteProductRequest{})

			Expect(err).NotTo(BeNil())

		})

	})

	Context("ViewProduct", func() {
		It("SUCCESS", func() {
			mockCarttU.EXPECT().ViewProduct(any).Return(
				[]entities.Carts{
					{
						ProductCode: "P001",
						ProductName: "Product 1",
						Quantity:    4,
					},
				},
				nil,
			)

			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			result, err := cartRpcClient.ViewProduct(ctxServiceToken, &pbCarts.Empty{})

			Expect(err).To(BeNil())
			Expect(result.GetProductsList()[0].GetProductCode()).To(Equal("P001"))
			Expect(result.GetProductsList()[0].GetProductName()).To(Equal("Product 1"))
			Expect(result.GetProductsList()[0].GetQuantity()).To(Equal(int32(4)))

		})

		It("Error", func() {
			mockCarttU.EXPECT().ViewProduct(any).Return(
				[]entities.Carts{},
				fmt.Errorf("Error"),
			)
			grpcService = NewCartsService(mockCarttU, zap.NewExample())
			cartRpcClient := callGrpcClient(grpcService)

			_, err := cartRpcClient.ViewProduct(ctxServiceToken, &pbCarts.Empty{})

			Expect(err).NotTo(BeNil())

		})

	})

})

func callGrpcClient(grpcServ pb.CartsServiceServer) pb.CartsServiceClient {
	logger := common.NewLogger()
	cJW := config.ConfigJWTEnv{
		JwtSecret:        "hkajhdjdkjdjadje249482kjdadadda",
		DurationTRefresh: 15 * time.Minute,
	}
	jwtManager := middleware.NewJWTManager(&cJW, logger)

	interceptor := interceptor.NewAuthInterceptor(jwtManager, logger)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterCartsServiceServer(grpcServer, grpcServ)

	listener, err := net.Listen("tcp", ":0") // random available port
	Expect(err).To(BeNil())

	go grpcServer.Serve(listener)
	return newTestQuotesClient(listener.Addr().String())
}

func newTestQuotesClient(serverAddress string) pb.CartsServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	Expect(err).To(BeNil())
	return pb.NewCartsServiceClient(conn)
}
