package main

import (
	"embed"
	"log"

	"boilerplate/server"
	"boilerplate/server/config"
	"boilerplate/server/middleware"
)

//go:embed all:frontend/webapp/build
var staticFiles embed.FS

func banner() {
	log.Println("Starting boilerplate server...")
}

func main() {
	banner()
	cfg := config.Load()

	// Create server and configure routes
	httpServer := server.New(cfg.Server, staticFiles, middleware.LoggingMiddleware, middleware.CORSMiddleware)
	httpServer.AddAPIRoutes() // All endpoints are under /api prefix
	httpServer.Start()
}
