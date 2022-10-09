package client

import (
	"fmt"
	"net/http"
)

// DoOps함수는 클라이언트를 매개변수로 받은 다음 google.com에 연결
func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of DoOps:", resp.StatusCode)

	return nil
}

// DefaultGetGolang함수는 golang.org 주소로 get요청을 하기 위한 클라이언트
func DefaultGetGolang() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}

	fmt.Println("results of DefaultGetGolang:",
		resp.StatusCode)

	return nil
}
