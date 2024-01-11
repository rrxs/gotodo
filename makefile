.PHONY: default run build test docs clean

APP_NAME=gotodo

default: run-n-docs

run:
	@go run main.go
run-n-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clear: 
	@rm -f $(APP_NAME)
	@rm -rf ./docs