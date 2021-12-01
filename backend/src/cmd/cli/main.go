package main

import (
	"backend/datastore"
	"backend/db"
	"backend/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func loggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	db := db.MustOpen("db.sqlite3")
	datastore, err := datastore.New(db)
	if err != nil {
		log.Fatalf("Failed to open Datastore: %v", err)
	}
	server := server.New(":9595")
	server.AddHandler("/api/memory", datastore.MemoryHandler)
	server.AddHandler("/", loggingHandler(http.FileServer(http.Dir("."))).ServeHTTP)

	onInterrupt(func() { server.Stop() })
	server.Start()
}

func onInterrupt(callback func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println(sig)
		callback()
	}()
}
