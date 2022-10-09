package client

import (
	"crypto/tls"
	"net/http"
)

// Setup 함수는 클라이어느를 구성하고 전역 DefaultClient를 재정의
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient

	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}

	if nop {
		c.Transport = &NopTransport{}
	}

	http.DefaultClient = c
	return c
}

// NopTransport 는 아무일도 하지 않는 No-operation 전송
type NopTransport struct {
}

// RoundTrip 함수는 RoundTripper 인터페이스를 구현
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	// 초기화되지 않은 응답
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
