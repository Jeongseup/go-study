package main

import (
	"fmt"
	"net/http"
)

// type Handler .. interface {}

type MyWebServerType bool

func (m MyWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "well, hello there!!")
}

func main() {
	var k MyWebServerType

	http.ListenAndServe("localhost:8000", k)
}
