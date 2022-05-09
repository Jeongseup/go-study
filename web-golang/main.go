package main

import (
	"net/http"

	"github.com/Jeongseup/go-study/web-golang/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
