package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	signalChan chan struct{}
	mux        *http.ServeMux
	server     *http.Server
}

func New(addr string) Server {
	mux := http.NewServeMux()
	server := http.Server{Addr: addr, Handler: mux}

	return Server{
		signalChan: make(chan struct{}),
		mux:        mux,
		server:     &server,
	}
}

func (server Server) AddHandler(pattern string, handler http.HandlerFunc) {
	server.mux.Handle(pattern, handler)
}

func (server Server) Start() {
	go func() {
		if err := server.server.ListenAndServe(); err == http.ErrServerClosed {
			log.Println("Server stopped")
			return
		} else {
			log.Fatal(err)
		}
	}()
	<-server.signalChan
	server.server.Shutdown(context.Background())
}

func (server Server) Stop() {
	close(server.signalChan)
}
