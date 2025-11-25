package main

import (
	"embed"
	"fmt"

	"boilerplate/backend/hello"
	"boilerplate/backend/server"
)

//go:embed all:build-frontend
var staticFiles embed.FS

func banner() {
	fmt.Println("Starting boilerplate server...")
}

func main() {
	banner()

	httpServer := server.New(3001, staticFiles, server.LoggingMiddleware, server.CORSMiddleware)
	httpServer.ApiEndpoint("GET /hello", hello.Handler)
	httpServer.Start()
}
