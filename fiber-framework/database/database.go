package database

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DatabaseInstance struct {
	DB *bun.DB
}

var Database DatabaseInstance

func ConnectDatabase() {
	dsn := "postgres://student:@localhost/study?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
		//  pgdriver.WithTLSConfig(...)
	))

	// TODO : 에러체크 필ㅛ
	db := bun.NewDB(pgdb, pgdialect.New())

	// if err != nil {
	// 	log.Fatal("Failed to connect to the database! \n", err.Error())
	// 	os.Exit(2)
	// }

	log.Println("Connected to the database successfully")

	// Bun debug hook
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	Database = DatabaseInstance{DB: db}
}
