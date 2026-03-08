.PHONY: run run-dev build

run:
	go run main.go

run-dev:
	DEV=1 go run main.go

build:
	npm run build
	go build -o bin/server main.go
