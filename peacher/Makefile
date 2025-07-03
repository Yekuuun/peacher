APP_NAME=peacher
BUILD_DIR=bin

.PHONY: all build run test clean vet tidy

all: build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server

run: build
	./$(BUILD_DIR)/$(APP_NAME)

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BUILD_DIR)
