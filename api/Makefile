DOCKER_IMAGE_NAME:=pr-leaderboard
DOCKER_IMAGE_REPO:=devteamtools/$(DOCKER_IMAGE_NAME)
DOCKER_IMAGE_SHA:=$(DOCKER_IMAGE_REPO):$(shell git rev-parse HEAD | cut -c 1-8)
DOCKER_IMAGE_LATEST:=$(DOCKER_IMAGE_REPO):LATEST
DOCKER_HUB_USERNAME:=devanarchy

.PHONY: all

all: clean get test build publish

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
	docker build -t $(DOCKER_IMAGE_NAME) .
	docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_SHA)
	docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_LATEST)

publish:
	docker login -u $(DOCKER_HUB_USERNAME) -p $(DOCKER_HUB_TOKEN)
	docker push $(DOCKER_IMAGE_SHA)
	docker push $(DOCKER_IMAGE_LATEST)
	docker logout