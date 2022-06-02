package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWolrd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func handleRequests() {
	// create new router
	myRouter := mux.NewRouter().StrictSlash(true)

	// router's handler func interfaces
	myRouter.HandleFunc("/", helloWolrd).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	// http listen with defined myRouter
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("")
	InitialMigration()
	handleRequests()
}
