BINARY_NAME=softbase

BASE_DIR=./examples/base

MAIN_DIR=$(BASE_DIR)/main.go

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_DIR)

run: build
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)

help:
	@echo "make build - build the binary"
	@echo "make run - run the binary"
	@echo "make clean - remove the binary"

.PHONY: all build run clean help
