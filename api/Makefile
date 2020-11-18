COMMIT_SHA:=$(shell git rev-parse HEAD | cut -c 1-8)

.PHONY: all

all: clean get test build

clean:
	rm -rf bin

get:
	go mod vendor

test:
	go test ./...

build: build-macos build-linux build-docker-image

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/app main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/app main.go

build-docker-image:
	docker build -t devteamtools/pr-leaderboard:$(COMMIT_SHA) .
	docker tag devteamtools/pr-leaderboard:$(COMMIT_SHA) devteamtools/pr-leaderboard:LATEST