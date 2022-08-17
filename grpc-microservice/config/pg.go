package config

import (
	"database/sql"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewPgConnection(config Config) *bun.DB {
	postgresUrl := config.Get("POSTGRES_URL")

	pg := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(postgresUrl)))
	db := bun.NewDB(pg, pgdialect.New())

	isDebug, _ := strconv.ParseBool(config.Get("DB_DEBUG"))
	if isDebug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	}

	return db
}
