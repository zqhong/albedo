# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=albedo

all: build

build: gotool
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)

run:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	./build/$(BINARY_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

gotool:
	gofmt -w .
