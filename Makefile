# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
WEB_BIN_NAME=albedo
CLI_BIN_NAME=albedo-cli

all: build

build: build-web build-cli

build-web: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME) -tags=jsoniter -v web.go

build-cli: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME) -tags=jsoniter -v cli.go

run-web: build-web
	./build/$(WEB_BIN_NAME)

run-cli: build-cli
	./build/$(CLI_BIN_NAME)

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(WEB_BIN_NAME)
	rm -f ./build/$(CLI_BIN_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

gotool:
	gofmt -w .
