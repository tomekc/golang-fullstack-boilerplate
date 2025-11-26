package main

import (
	"embed"
	"log"

	"boilerplate/backend/config"
	"boilerplate/backend/hello"
	"boilerplate/backend/server"
)

//go:embed all:build-frontend
var staticFiles embed.FS

func banner() {
	log.Println("Starting boilerplate server...")
}

func main() {
	banner()
	cfg := config.Load()

	httpServer := server.New(cfg.Server, staticFiles, server.LoggingMiddleware, server.CORSMiddleware)
	httpServer.ApiEndpoint("GET /hello", hello.Handler)
	httpServer.Start()
}
