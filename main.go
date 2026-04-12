package main

import (
	"log"

	"boilerplate/server"
	"boilerplate/server/config"
	"boilerplate/server/middleware"
)

func banner() {
	log.Println("Starting boilerplate server...")
}

func main() {
	banner()
	cfg := config.Load()

	// Create server and configure routes
	httpServer := server.New(cfg.Server)
	httpServer.Use(middleware.Logging)
	httpServer.Use(middleware.CORS)
	httpServer.RegisterAPIRoutes() // All endpoints are under /api prefix. Inject any dependencies here.
	httpServer.Start()
}
