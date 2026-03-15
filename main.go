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
	httpServer := server.New(cfg.Server, staticFiles)
	httpServer.Use(middleware.LoggingMiddleware)
	httpServer.Use(middleware.CORSMiddleware)
	httpServer.RegisterAPIRoutes() // All endpoints are under /api prefix. Inject any dependencies here.
	httpServer.Start()
}
