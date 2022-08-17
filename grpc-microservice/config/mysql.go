package config

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewMySQLConnection(envconfig Config) *bun.DB {
	dsn := envconfig.Get("MYSQL_URL") //"root:pass@/test"

	sqldb, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	isDebug, _ := strconv.ParseBool(envconfig.Get("DB_DEBUG"))
	if isDebug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	}

	return db
}
