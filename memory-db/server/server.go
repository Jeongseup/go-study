package server

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	listener         net.Listener
	quit             chan struct{}
	exited           chan struct{}
	db               memoryDB
	connections      map[int]net.Conn
	connCloseTimeout time.Duration
}

func NewServer() *Server {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to create listener", err.Error())
	}

	srv := &Server{
		listener:         l,
		quit:             make(chan struct{}),
		exited:           make(chan struct{}),
		db:               newDB(),
		connections:      map[int]net.Conn{},
		connCloseTimeout: 10 * time.Second,
	}

	go srv.serve()
	return srv
}

func (s *Server) serve() {
	fmt.Println("listening for clients")
	for {
		select {
		case <-srv.quit:
			fmt.Println("shutting downt the server")
		default:
			tcpListenter := srv.listener.(*net.TCPListener)
			err := toc
		}
	}
}

func (s *Server) Stop() {
	fmt.Println("stopping server")
}
