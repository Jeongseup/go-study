package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-study/3_go-bookstore/pkg/config"
)

var db *pg.DB

type Book struct {
	pg.DB
	Name        string `pg:""json:"name"`
	Author      string `json:"author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
