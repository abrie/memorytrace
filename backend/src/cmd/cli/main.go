package main

import (
	"backend/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

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
