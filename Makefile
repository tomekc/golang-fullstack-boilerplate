.PHONY: run run-dev build templ-generate templ-watch run-templ

run:
	go run main.go

run-dev:
	DEV=1 go run main.go

build:
	npm run build
	templ generate
	go build -o bin/server main.go

templ-generate:
	templ generate

templ-watch:
	templ generate --watch --proxy="http://localhost:3001" --cmd="DEV=1 go run main.go"

run-templ:
	templ generate && DEV=1 go run main.go
