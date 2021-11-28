package main

import (
	"backend/datastore"
	"backend/db"
	"backend/server"
	"log"
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
	log.Println(datastore)
	server := server.New(":9595")

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
