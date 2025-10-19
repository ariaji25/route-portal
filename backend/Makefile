.PHONY: build run test unit-test clean

WORKDIR := $(shell pwd)/backend
BINARY_NAME := main
BUILD_DIR := bin

prep: 
	@echo "Preparing backend..."
	rm -rf .data
	mkdir -p .data
	touch .data/routes.yaml
	make clean && go mod tidy

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -failfast -v -p 1 -coverprofile cover.out ./test -coverpkg=./...
	go tool cover -html cover.out -o cover.html

clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)
	rm -f cover.out cover.html
