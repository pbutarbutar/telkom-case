package carts

import (
	"context"
	"fmt"
	"testing"

	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"grpc-microservice/shared/models/entities"
	mock_repo "grpc-microservice/shared/repository/carts/mocks"

	"grpc-microservice/shared/utils"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Carts Suite")
}

var _ = Describe("Checking Carts Usecase", Ordered, func() {
	var (
		uc       CartsUseCaseImpl
		ctrl     *gomock.Controller
		mockCart *mock_repo.MockCartsRepository
		utl      utils.Utility
	)
	any := gomock.Any()

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockCart = mock_repo.NewMockCartsRepository(ctrl)
		uc = CartsUseCaseImpl{
			Logger: zap.NewExample(),
			Repo:   mockCart,
			Utils:  utils.Utility{},
		}

	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("NewCommentUsecase", func() {
		It("SUCCESS", func() {
			cartsUC := NewCartsUseCaseImpl(mockCart, utl, zap.NewExample())
			Expect(cartsUC).NotTo(BeNil())
		})
	})

	Context("AddProduct", func() {
		It("SUCCESS New Cart", func() {

			mockCart.EXPECT().IsExistProductByCode(any, any).Return(
				false,
				&entities.Carts{},
				nil,
			)

			mockCart.EXPECT().AddProduct(any, any).Return(
				entities.Carts{
					ProductCode: "P001",
					ProductName: "Produk 1",
					Quantity:    2,
				},
				nil,
			)

			res, err := uc.AddProduct(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Produk 1",
				Quantity:    2,
			})

			Expect(err).To(BeNil())
			Expect(res.ProductCode).To(Equal("P001"))
			Expect(res.Quantity).To(Equal(int32(2)))

		})

		It("SUCCESS Product in Cart is exist", func() {

			mockCart.EXPECT().IsExistProductByCode(any, any).Return(
				true,
				&entities.Carts{
					ProductCode: "P001",
					ProductName: "Produk 1",
					Quantity:    2,
				},
				nil,
			)

			mockCart.EXPECT().UpdateCart(any, any).Return(
				entities.Carts{
					ProductCode: "P001",
					ProductName: "Produk 1",
					Quantity:    4,
				},
				nil,
			)

			res, err := uc.AddProduct(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Produk 1",
				Quantity:    2,
			})

			Expect(err).To(BeNil())
			Expect(res.ProductCode).To(Equal("P001"))
			Expect(res.Quantity).To(Equal(int32(4)))

		})

		It("ERROR", func() {

			mockCart.EXPECT().IsExistProductByCode(any, any).Return(
				true,
				&entities.Carts{},
				fmt.Errorf("error"),
			)

			_, err := uc.AddProduct(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Produk 1",
				Quantity:    2,
			})
			Expect(err).To(Equal(fmt.Errorf("error")))

		})
	})

	Context("DeleteProduct", func() {
		It("SUCCESS", func() {
			mockCart.EXPECT().DeleteProduct(any, any).Return(
				true,
				nil,
			)
			res, err := uc.DeleteProduct(context.Background(), "P001")
			Expect(err).To(BeNil())
			Expect(res).To(BeTrue())

		})

		It("SUCCESS is false", func() {
			mockCart.EXPECT().DeleteProduct(any, any).Return(
				false,
				nil,
			)
			res, err := uc.DeleteProduct(context.Background(), "P001")
			Expect(err).To(BeNil())
			Expect(res).To(BeFalse())

		})

		It("ERROR", func() {
			mockCart.EXPECT().DeleteProduct(any, any).Return(
				true,
				fmt.Errorf("error"),
			)
			_, err := uc.DeleteProduct(context.Background(), "")
			Expect(err).To(Equal(fmt.Errorf("error")))

		})
	})

	Context("ViewProduct", func() {
		It("SUCCESS", func() {
			mockCart.EXPECT().ViewProducts(any).Return(
				[]entities.Carts{
					{
						ProductCode: "P001",
						ProductName: "Produk 1",
						Quantity:    4,
					},
				},
				nil,
			)
			res, err := uc.ViewProduct(context.Background())
			Expect(err).To(BeNil())
			Expect(len(res)).To(Equal(1))
			Expect(res[0].ProductCode).To(Equal("P001"))
			Expect(res[0].Quantity).To(Equal(int32(4)))

		})

		It("ERROR", func() {
			mockCart.EXPECT().ViewProducts(any).Return(
				[]entities.Carts{},
				fmt.Errorf("error"),
			)
			_, err := uc.ViewProduct(context.Background())
			Expect(err).To(Equal(fmt.Errorf("error")))

		})
	})

})
