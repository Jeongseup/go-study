package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func startBadTestHTTPServerV1() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(60 * time.Second) // <- 여기서 병목
		fmt.Fprint(w, "Hello World")
	}))
	return ts
}

// / 이거 테스트 돌려도 서버쪽 세션이 종료가 안되어서 테스트가 이상해짐..
func TestFetchBadRemoteResourceV1(t *testing.T) {
	ts := startBadTestHTTPServerV1()
	defer ts.Close()

	client := createHTTPClientWithTimeout(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, ts.URL)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("Expected error to contain: context deadline exceeded, Got: %v", err.Error())
	}
}
