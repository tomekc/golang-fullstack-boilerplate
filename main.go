package main

import (
	"fmt"

	"boilerplate/backend/hello"
	"boilerplate/backend/server"
)

func banner() {
	fmt.Println("Starting boilerplate server...")
}

func main() {
	banner()

	httpServer := server.New(3001, server.LoggingMiddleware, server.CORSMiddleware)
	httpServer.ApiEndpoint("GET /hello", hello.Handler)
	httpServer.Start()
}
