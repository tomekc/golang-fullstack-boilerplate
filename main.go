package main

import (
	"log"

	"boilerplate/server"
	"boilerplate/server/config"
)

func banner() {
	log.Println("Starting boilerplate server...")
}

func main() {
	banner()
	cfg := config.Load()

	httpServer := server.New(cfg.Server)
	httpServer.Start()
}
