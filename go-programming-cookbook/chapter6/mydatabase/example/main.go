package main

import (
	"mydatabase"

	_ "github.com/go-pg/pg"
)

func main() {
	db, err := mydatabase.Setup()
	if err != nil {
		panic(err)
	}

	if err := mydatabase.Exec(db); err != nil {
		panic(err)
	}
}
