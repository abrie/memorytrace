package main

import (
	"backend/datastore"
	"backend/db"
	"backend/models/memory"
	"backend/server"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	db := db.MustOpen("file:db.sqlite3")
	datastore, err := datastore.New(db)
	if err != nil {
		log.Fatalf("Failed to open Datastore: %v", err)
	}
	server := server.New(":9595")
	server.AddHandler("/api/memory", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var m memory.Memory
			if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
				log.Printf("Bad Request: %v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := datastore.PutMemory(m); err != nil {
				log.Printf("Internal Server Error: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

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
