package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	signalChan chan struct{}
	handler    *http.ServeMux
	server     *http.Server
}

func New(addr string) Server {
	handler := http.NewServeMux()
	server := http.Server{Addr: addr, Handler: handler}

	return Server{
		signalChan: make(chan struct{}),
		handler:    handler,
		server:     &server,
	}
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
