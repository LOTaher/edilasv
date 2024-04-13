BINARY_NAME=softbase

CMD_DIR=./cmd/$(BINARY_NAME)

all: build

build:
	go build -o $(BINARY_NAME) $(CMD_DIR)

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
