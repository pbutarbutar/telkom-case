package carts

import (
	"context"
	"database/sql"
	"fmt"
	pbCarts "grpc-microservice/grpc-app/protos/carts"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cart Repo Suite")
}

var _ = Describe("Checking Cart Repository", Ordered, func() {
	var cartRepo CartsRepository
	var dbMock *sql.DB
	var bunDB *bun.DB
	var mockDB sqlmock.Sqlmock

	//pgDialect := pgdialect.New()
	mySQLDialect := mysqldialect.New()

	BeforeEach(func() {
		dbMock, mockDB, _ = sqlmock.New()
		bunDB = bun.NewDB(dbMock, mySQLDialect)
		cartRepo = NewCartsRepositoryImpl(bunDB, zap.NewExample())
	})
	AfterEach(func() {
		dbMock.Close()
	})

	Context("IsExistProductByCode", func() {
		It("SUCCESS", func() {
			rows := sqlmock.NewRows([]string{"product_code"}).AddRow(true)
			mockDB.ExpectQuery("SELECT (.+)").WillReturnRows(rows)

			res, carts, err := cartRepo.IsExistProductByCode(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    2,
			},
			)

			Expect(err).To(BeNil())
			Expect(res).To(BeTrue())
			Expect(carts.ProductCode).To(Equal("P001"))

		})

		It("Not Found", func() {
			rows := sqlmock.NewRows([]string{"product_code"}).AddRow(false)
			mockDB.ExpectQuery("SELECT (.+)").WillReturnRows(rows)

			res, _, err := cartRepo.IsExistProductByCode(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    2,
			},
			)

			Expect(err).To(BeNil())
			Expect(res).To(BeFalse())

		})

		It("FAIL", func() {
			mockDB.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))

			res, _, err := cartRepo.IsExistProductByCode(context.Background(), &pbCarts.AddProductRequest{})

			Expect(err).NotTo(BeNil())
			Expect(res).To(BeFalse())
		})

	})

	Context("AddProduct", func() {
		It("SUCCESS", func() {
			mockDB.ExpectExec("INSERT (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := cartRepo.AddProduct(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    2,
			},
			)

			Expect(err).To(BeNil())
			Expect(res.ProductCode).To(Equal("P001"))
			Expect(res.ProductName).To(Equal("Product 1"))
			Expect(res.Quantity).To(Equal(int32(2)))

		})

		It("FAIL", func() {
			mockDB.ExpectExec("INSERT (.+)").WillReturnError(fmt.Errorf("error"))

			_, err := cartRepo.AddProduct(context.Background(), &pbCarts.AddProductRequest{})

			Expect(err).NotTo(BeNil())
		})

	})

	Context("UpdateCart", func() {
		It("SUCCESS", func() {
			mockDB.ExpectExec("UPDATE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := cartRepo.UpdateCart(context.Background(), &pbCarts.AddProductRequest{
				ProductCode: "P001",
				ProductName: "Product 1",
				Quantity:    2,
			},
			)

			Expect(err).To(BeNil())
			Expect(res.ProductCode).To(Equal("P001"))
			Expect(res.ProductName).To(Equal("Product 1"))
			Expect(res.Quantity).To(Equal(int32(2)))

		})

		It("FAIL", func() {
			mockDB.ExpectExec("UPDATE (.+)").WillReturnError(fmt.Errorf("error"))

			_, err := cartRepo.UpdateCart(context.Background(), &pbCarts.AddProductRequest{})

			Expect(err).NotTo(BeNil())
		})

	})

	Context("DeleteProduct", func() {
		It("SUCCESS", func() {
			mockDB.ExpectExec("DELETE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

			res, err := cartRepo.DeleteProduct(context.Background(), "P001")

			Expect(err).To(BeNil())
			Expect(res).To(BeTrue())

		})

		It("FAIL", func() {
			mockDB.ExpectExec("DELETE (.+)").WillReturnError(fmt.Errorf("error"))

			_, err := cartRepo.DeleteProduct(context.Background(), "")

			Expect(err).NotTo(BeNil())
		})

	})

	Context("ViewProduct", func() {
		It("SUCCESS", func() {
			result := sqlmock.NewRows([]string{"product_code", "product_name", "quantity"}).
				AddRow("P001", "Product 1", "2")
			mockDB.ExpectQuery("SELECT (.+)").WillReturnRows(result)

			res, err := cartRepo.ViewProducts(context.Background())

			Expect(err).To(BeNil())

			Expect(res[0].ProductCode).To(Equal("P001"))
			Expect(res[0].ProductName).To(Equal("Product 1"))
			Expect(res[0].Quantity).To(Equal(int32(2)))
		})

		It("Not Found", func() {
			rows := sqlmock.NewRows([]string{"product_code"}).AddRow("")
			mockDB.ExpectQuery("SELECT (.+)").WillReturnRows(rows)

			res, err := cartRepo.ViewProducts(context.Background())
			Expect(err).To(BeNil())
			Expect(res[0].ProductCode).To(Equal(""))
		})

		It("FAIL", func() {
			mockDB.ExpectQuery("SELECT (.+)").WillReturnError(fmt.Errorf("error"))

			_, err := cartRepo.ViewProducts(context.Background())

			Expect(err).NotTo(BeNil())
		})

	})

})
