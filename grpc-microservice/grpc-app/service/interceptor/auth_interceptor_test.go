package interceptor

import (
	"grpc-microservice/common"
	"grpc-microservice/middleware"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
)

//Verify assigning mock
var dataVerify func(accessToken string) (*middleware.UserClaims, error)
var dataGenerate func(user *middleware.User) (string, error)

type JwtManagerMock struct{}

func (j JwtManagerMock) Verify(accessToken string) (*middleware.UserClaims, error) {
	return dataVerify(accessToken)
}
func (j JwtManagerMock) Generate(user *middleware.User) (string, error) {
	return dataGenerate(user)
}

var _ = Describe("Interceptor", func() {

	BeforeEach(func() {

	})
	Describe("Test Unary", func() {
		It("should can access Unary", func() {
			mockJM := JwtManagerMock{}
			dataVerify = func(accessToken string) (*middleware.UserClaims, error) {
				return &middleware.UserClaims{
					UserID:  "XXXXX",
					HumanID: "YYYYY",
				}, nil
			}
			dataGenerate = func(user *middleware.User) (string, error) {
				return "XXXXX", nil
			}
			logger := common.NewLogger()
			newIntecptor := NewAuthInterceptor(mockJM, logger)

			xG := newIntecptor.Unary()
			Expect(xG).NotTo(BeNil())
		})
	})

	Describe("Test Stream", func() {
		It("should can access Stream", func() {
			mockJM := JwtManagerMock{}
			dataVerify = func(accessToken string) (*middleware.UserClaims, error) {
				return &middleware.UserClaims{
					UserID:  "XXXXX",
					HumanID: "YYYYY",
				}, nil
			}
			dataGenerate = func(user *middleware.User) (string, error) {
				return "XXXXX", nil
			}
			logger := common.NewLogger()
			newIntecptor := NewAuthInterceptor(mockJM, logger)

			xG := newIntecptor.Stream()

			Expect(xG).NotTo(BeNil())
		})
	})
})
