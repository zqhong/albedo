# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=albedo

all: build

build: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

run: build
	./build/$(BINARY_NAME)

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

gotool:
	gofmt -w .
