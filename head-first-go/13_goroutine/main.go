package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://example.com")
	go responseSize("https://golang.com")
	go responseSize("https://golang.org/doc")

	time.Sleep(5 * time.Second)

}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// main함수가 종료될 때 네트워크 연결을 해제
	defer response.Body.Close()
	// 모든 응답 데이터를 읽는다
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// byte slice의 크기는 페이지의 크기가 같음
	fmt.Println(len(body))
}
