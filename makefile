# Simple Makefile for a Go project
DIR_SRC=./cmd/app
BIN=bin/app

wire:
	wire ./...

test:
	go test -cover ./...

run: wire
	go run $(DIR_SRC)

build: wire
	go build -o $(BIN) $(DIR_SRC)

gen-doc:
	swag fmt -g ./internal/deliver/http/server/http.go
	swag init -g ./internal/deliver/http/server/http.go -o document/swagger/
