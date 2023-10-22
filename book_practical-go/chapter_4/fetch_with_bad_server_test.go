package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func startBadTestHTTPServerV2(shutdownServer chan struct{}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-shutdownServer
		fmt.Fprint(w, "Hello World")
	}))
	return ts
}

func TestFetchBadRemoteResourceV2(t *testing.T) {
	shutdownServer := make(chan struct{})
	ts := startBadTestHTTPServerV2(shutdownServer)
	defer ts.Close()
	// 아래 테스트가 다 돌아서 서버쪽을 종료시키려고 하면 이렇게 shutdown signal을 dummy sturct{}로 보내줌
	defer func() {
		shutdownServer <- struct{}{}
	}()

	client := createHTTPClientWithTimeout(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, ts.URL)
	if err == nil {
		t.Fatal("Expected non-nil error")
	}

	if !strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers") {
		t.Fatalf("Expected error to contain: Client.Timeout exceeded while awaiting headers, Got: %v", err.Error())
	}
}
