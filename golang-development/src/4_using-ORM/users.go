package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed")
	}

	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All users Endpoint hit")

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to db")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "create new user successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get data in query
	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "delete a user successfully")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get data in query
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Successfully update a user info")
}
