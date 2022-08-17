package middleware

import (
	"fmt"
	"grpc-microservice/common"
	"grpc-microservice/config"
	"strings"
	"time"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Middleware JWT", func() {

	BeforeEach(func() {

	})

	Describe("Test JWT Middleware", func() {
		It("Generate", func() {
			cJW := config.ConfigJWTEnv{
				JwtSecret:        "XZccxxcxcxcxc",
				DurationTRefresh: 15 * time.Minute,
			}
			logger := common.NewLogger()
			jwtManager := NewJWTManager(&cJW, logger)
			usr := User{
				UserID:  "XXXXX",
				HumanID: "XXXXX",
				Email:   "XXXX@gmail.com",
			}
			res, err := jwtManager.Generate(&usr)

			validTokenChar := strings.Split(res, ".")
			Expect(len(validTokenChar)).Should(Equal(3))
			Expect(err).To(BeNil())
		})
	})

	Describe("Test JWT Middleware", func() {
		It("Verify", func() {
			cJW := config.ConfigJWTEnv{
				JwtSecret:        "XZccxxcxcxcxc",
				DurationTRefresh: 15 * time.Minute,
			}
			logger := common.NewLogger()
			jwtManager := NewJWTManager(&cJW, logger)
			usr := User{
				UserID:  "XXXXX",
				HumanID: "XXXXX",
				Email:   "XXXX@gmail.com",
			}
			tokenS, _ := jwtManager.Generate(&usr)
			res, err := jwtManager.Verify(tokenS)
			fmt.Println(res)
			Expect(res.UserID).To(Equal("XXXXX"))
			Expect(err).To(BeNil())
		})
	})
})
