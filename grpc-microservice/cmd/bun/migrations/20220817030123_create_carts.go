package migrations

import (
	"context"
	"fmt"
	"grpc-microservice/shared/models/entities"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		if _, err := db.NewCreateTable().
			Model((*entities.Carts)(nil)).
			ModelTableExpr("carts").
			IfNotExists().
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		if _, err := db.NewCreateTable().
			Model((*entities.Carts)(nil)).
			ModelTableExpr("carts").
			IfNotExists().
			Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
