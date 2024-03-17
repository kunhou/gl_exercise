# Simple Makefile for a Go project
DIR_SRC=./cmd/app
BIN=bin/app

run:
	go run $(DIR_SRC)

build:
	go build -o $(BIN) $(DIR_SRC)