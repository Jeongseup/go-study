package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}

func createHTTPGetRequestWithTrace(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		TLSHandshakeStart: func() {
			fmt.Printf("TLS HandShake Start\n")
		},
		TLSHandshakeDone: func(connState tls.ConnectionState, err error) {
			fmt.Printf("TLS HandShake Done\n")
		},

		PutIdleConn: func(err error) {
			fmt.Printf("Put Idle Conn Error: %+v\n", err)
		},
	}
	ctxTrace := httptrace.WithClientTrace(req.Context(), trace)
	req = req.WithContext(ctxTrace)
	return req, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from")
		os.Exit(1)
	}

	d := 5 * time.Second
	ctx := context.Background()
	client := createHTTPClientWithTimeout(d)

	// request once
	// body, err := fetchRemoteResource(client, os.Args[1])
	// if err != nil {
	// 	fmt.Fprintf(os.Stdout, "%#v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Fprintf(os.Stdout, "%s\n", body)

	// request loop
	req, err := createHTTPGetRequestWithTrace(ctx, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, _ := client.Do(req)
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		fmt.Printf("Resp protocol: %#v\n", resp.Proto)
		time.Sleep(1 * time.Second)
		fmt.Println("--------")
	}
}
