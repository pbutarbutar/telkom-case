package health

import (
	"context"
	"database/sql"
	"fmt"
	pb "grpc-microservice/grpc-app/protos"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"go.uber.org/zap"
)

var _ = Describe("Orders Go GRGP Services", func() {
	var hR pb.HealthServer
	var ctx context.Context

	BeforeEach(func() {
	})

	It("should be status is SERVING", func() {
		sqldb, _ := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
		sqldb.SetMaxOpenConns(1)
		DB = bun.NewDB(sqldb, sqlitedialect.New())
		DeferCleanup(DB.Close)

		hR = NewHalthServer(DB, zap.NewExample())
		result, err := hR.Check(ctx, &pb.HealthCheckRequest{Service: "Test"})
		Expect(err).To(BeNil())
		strResult := pb.HealthCheckResponse_ServingStatus(result.GetStatus())
		strResultVal := fmt.Sprint(strResult)
		Expect(strResultVal).To(Equal("SERVING"))
	})

	It("should be status is NOT_SERVING DB can't PING", func() {
		sqldb, _ := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
		sqldb.SetMaxOpenConns(1)
		DB = bun.NewDB(sqldb, sqlitedialect.New())
		DB.Close()

		hR = NewHalthServer(DB, zap.NewExample())
		result, err := hR.Check(ctx, &pb.HealthCheckRequest{Service: "aaa"})
		Expect(err).NotTo(BeNil())
		strResult := pb.HealthCheckResponse_ServingStatus(result.GetStatus())
		strResultVal := fmt.Sprint(strResult)
		Expect(strResultVal).To(Equal("NOT_SERVING"))
	})

})
