package routes

import (
	"3_go-bookstore/pkg/bcontrollers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", bcontrollers.CreateBook).Methods("POST")
	router.HandleFunc("/book", bcontrollers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", bcontrollers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", bcontrollers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bcontrollers.DeleteBook).Methods("DELETE")
}
