package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/slok/reload"
)

const (
	reloadKeySIGHUP   = "signal-sighup"
	reloadKeySIGINT   = "signal-sigint"
	reloadKeySIGTERM  = "signal-sigterm"
	reloadKeyWebhhok1 = "signal-wh-1"
	reloadKeyWebhhok2 = "signal-wh-2"
)

func main() {

	reloadSvc := reload.NewManager()

	// Add reloaders.
	reloadSvc.Add(0, reload.ReloaderFunc(func(ctx context.Context, id string) error {
		fmt.Printf("reloader 1: %s\n", id)
		return nil
	}))

	// HTTP webhook reloader.
	{
		c := make(chan string)
		router := http.NewServeMux()
		router.Handle("/w1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { c <- reloadKeyWebhhok1 }))
		router.Handle("/w2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { c <- reloadKeyWebhhok2 }))

		addr := ":8080"
		fmt.Printf("webhook '/w1' and '/w2' listening at %s...\n", addr)
		go func() {
			err := http.ListenAndServe(addr, router)
			if err != nil {
				panic(err)
			}
		}()

		reloadSvc.On(reload.NotifierFunc(func(ctx context.Context) (string, error) {
			return <-c, nil
		}))
	}

	err := reloadSvc.Run(context.Background())
	if err != nil {
		log.Panic(err.Error())
	}
}
