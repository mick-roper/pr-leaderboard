.PHONY: all

all: clean build

clean:
	rm -rf bin

build:
	go build -o bin/pr-leaderboard app/main.go && cp app/index.html bin/