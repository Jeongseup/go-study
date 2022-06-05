package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("New connection")

		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})


	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	http.Handle("/socket.io/", server)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
