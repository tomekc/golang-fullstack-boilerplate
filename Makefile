.PHONY: help run run-dev build templ-generate templ-watch run-templ

.DEFAULT_GOAL := help

help: ## Show this help message
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} \
		/^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

run: ## Run the server with 'go run'
	go run main.go

build: ## Generate Templ code and build production binary to bin/server
	templ generate
	go build -o bin/server main.go

templ-generate: ## Generate Go code from .templ files
	templ generate

templ-watch: ## Watch .templ files and auto-reload (proxies localhost:3000)
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run main.go"

run-templ: ## Generate Templ code and run server in DEV mode
	templ generate && DEV=1 go run main.go
