# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
WEB_BIN_NAME=albedo
CLI_BIN_NAME=albedo-cli

all: build

build: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME) -tags=jsoniter -v web.go
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME) -tags=jsoniter -v cli.go

run: build
	./build/$(WEB_BIN_NAME)

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(WEB_BIN_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

gotool:
	gofmt -w .
