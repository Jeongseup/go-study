package rest

import "net/http"

// APITransport는 모든 요청에 대해 SetBasicAuth를 수행
type APITransport struct {
	*http.Transport
	username, password string
}

// RoundTrip 함수는 기본전송을 수행하기 전에 기본적 인증을 진행
