package myapp

import (
	"fmt"
	"net/http"
)

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	fmt.Println(mux)

	return mux
}
