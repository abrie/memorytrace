package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	addr       string
	signalChan chan struct{}
	mux        *http.ServeMux
	server     *http.Server
}

func New(addr string) Server {
	mux := http.NewServeMux()
	server := http.Server{Addr: addr, Handler: mux}

	return Server{
		addr:       addr,
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
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	log.Printf(`Starting server at address "%s"`, server.addr)
	<-server.signalChan
	server.server.Shutdown(context.Background())
}

func (server Server) Stop() {
	close(server.signalChan)
}
