package config

import (
	"github.com/go-pg/pg/v10"
)

var (
	db *pg.DB
)

func Connect() {
	opt, err := pg.ParseURL("postgres://student:@localhost:5432/study")
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)
}

func GetDB() *pg.DB {
	return db
}
