package main

import (
	"fmt"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponsseWriter, *Request)
}

*/

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!s")
}

func main() {
	// 핸들러를 등록
	// '/'로 request가 들어왔을 때 어떻게 할 지?
	// node랑 다르게 req, res가 아니라 res, req라고 보면 될 듯?
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// f(출력) print
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello Bar Page</h1>")
	})

	// 위와 달리 Handler를 함수 형태가 아닌 인터페이스 형태로 제공
	// 기초 강좌에서 인터페이스 부분을 참고
	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
