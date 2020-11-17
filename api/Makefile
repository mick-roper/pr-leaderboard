.PHONY: all

all: clean get test build

clean:
	rm -rf bin

get:
	go mod vendor

test:
	go test ./...

build: build-macos build-linux

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/app main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/app main.go