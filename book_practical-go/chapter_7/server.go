package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("protocol=%s path=%s method=%s duration=%f", r.Proto, r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds())
	})
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("ping: Got a request")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "pong")
}

func doSomeWork(data []byte) {
	log.Println("body data: ", string(data))
	time.Sleep(1 * time.Second)
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	log.Println("I started processing the request")

	req, err := http.NewRequestWithContext(
		r.Context(),
		"GET",
		"http://localhost:8080/ping", nil,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	log.Println("Outgoing HTTP request")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	log.Println("Processing the response i got")

	go func() {
		doSomeWork(data)
		done <- true
	}()

	select {
	case <-done:
		log.Println("doSomeWork done: Continuing request processing")
	case <-r.Context().Done():
		log.Printf("Aborting request processing: %v\n", r.Context().Err())
		return
	}

	fmt.Fprint(w, string(data))
	log.Println("I finished processing the request")
}

func shutDown(ctx context.Context, s *http.Server, waitForShutdownCompletion chan struct{}) {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigch
	log.Printf("Got signal: %v . Server shutting down.", sig)

	childCtx, cancel := context.WithTimeout(
		ctx, 30*time.Second,
	)
	defer cancel()

	if err := s.Shutdown(childCtx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
	waitForShutdownCompletion <- struct{}{}
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	// for tls
	tlsCertFile := os.Getenv("TLS_CERT_FILE_PATH")
	tlsKeyFile := os.Getenv("TLS_KEY_FILE_PATH")

	if len(tlsCertFile) == 0 || len(tlsKeyFile) == 0 {
		log.Fatal("TLS_CERT_FILE_PATH and TLS_KEY_FILE_PATH must both be specified")
	}

	timeoutDuration := 30 * time.Second

	// for grace shutdown
	waitForShutdownCompletion := make(chan struct{})

	userHandler := http.HandlerFunc(handleUserAPI)
	hTimeout := http.TimeoutHandler(
		userHandler,
		timeoutDuration,
		"I ran out of time",
	)

	mux := http.NewServeMux()
	mux.Handle("/api/users/", hTimeout)
	mux.HandleFunc("/ping", handlePing)
	mux.HandleFunc("/api", apiHandler)

	m := loggingMiddleware(mux)

	s := http.Server{
		Addr:         listenAddr,
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go shutDown(context.Background(), &s, waitForShutdownCompletion)
	err := s.ListenAndServeTLS(tlsCertFile, tlsKeyFile)
	log.Print(
		"Waiting for shutdown to complete..",
	)
	<-waitForShutdownCompletion

	log.Fatal(err)
}
