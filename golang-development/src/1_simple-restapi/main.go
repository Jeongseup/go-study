package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "title1", Desc: "descrption1", Content: "Hello World"},
	}

	fmt.Println("hit1")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test post endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME PAGE ENDPOINT HIT")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticle).Methods("POST")

	myRouter.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequests()
}
