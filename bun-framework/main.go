package main

import (
	"context"
	"database/sql"
	"fmt"

	"example.com/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	ctx := context.Background()

	sqlite, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	sqlite.SetMaxOpenConns(1)

	db := bun.NewDB(sqlite, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	// res, err := db.ExecContext(ctx, "SELECT 1")
	// res, err := db.NewSelect().ColumnExpr("1").Exec(ctx)

	var num int
	// err = db.QueryRowContext(ctx, "SELECT 1").Scan(&num)
	db.NewSelect().ColumnExpr("1").Scan(ctx, &num)
	fmt.Println(num)

	db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)

	// Insert
	user := &models.User{Name: "admin"}
	_, err = db.NewInsert().Model(user).Exec(ctx)

	selectedUser := new(models.User)
	err = db.NewSelect().Model(selectedUser).Where("id = ?", 1).Scan(ctx, &selectedUser)
	fmt.Println(selectedUser)

	db.NewDropTable().Model((*models.User)(nil)).Exec(ctx)
}
