VERSION ?= $(shell git describe --tags)

IMAGE = yieldr/aws-cloudwatch-daemon
PKG = github.com/yieldr/aws-cloudwatch-daemon

OS ?= darwin
ARCH ?= amd64

GOBUILDFLAGS = -a -tags netgo -ldflags '-w'

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/aws-cloudwatch-daemon $(GOBUILDFLAGS)

test:
	go test

docker-all: docker-build docker-image docker-push

docker-build:
	@docker run -i --rm -v "$(PWD):/go/src/$(PKG)" $(IMAGE):build make build OS=linux

docker-test:
	@docker run -i --rm -v "$(PWD):/go/src/$(PKG)" $(IMAGE):build make test

docker-image:
	@docker build -t $(IMAGE):$(VERSION) .
	@docker tag -f $(IMAGE):$(VERSION) $(IMAGE):latest
	@echo " ---> $(IMAGE):$(VERSION)\n ---> $(IMAGE):latest"

docker-push:
	@docker push $(IMAGE):$(VERSION)
	@docker push $(IMAGE):latest

docker-builder-image:
	@docker build -t $(IMAGE):build -f Dockerfile.build .
