# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
WEB_BIN_NAME=web
CLI_BIN_NAME=cli
INSTALLER_BIN_NAME=installer

all: build

build: build-web build-cli

build-web: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME) -tags=jsoniter -v web.go

build-cli: gotool
	$(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME) -v cli.go

build-installer: gotool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(INSTALLER_BIN_NAME)-linux -v installer.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(INSTALLER_BIN_NAME)-win.exe -v installer.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(INSTALLER_BIN_NAME)-darwin -v installer.go

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
