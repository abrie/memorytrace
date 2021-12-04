package main

import (
	"backend/datastore"
	"backend/db"
	"backend/server"
	"backend/static"
	"flag"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
)

func main() {
	addr := flag.String("addr", ":80", "Address for server")
	dataPath := flag.String("data", "", "Path to datastore")
	staticPath := flag.String("static", "", "Path to serve as static files.")
	flag.Parse()

	absDataPath, err := filepath.Abs(*dataPath)
	if err != nil {
		log.Fatal(err)
	}

	absStaticPath, err := filepath.Abs(*staticPath)
	if err != nil {
		log.Fatal(err)
	}

	dbPath := path.Join(absDataPath, "db.sqlite3")
	db, err := db.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	datastore, err := datastore.Open(db)
	if err != nil {
		log.Fatal(err)
	}

	static := static.New(absStaticPath)

	server := server.New(*addr)
	server.AddHandler("/api/memory", datastore.MemoryHandler)
	server.AddHandler("/", static.Handler)

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
