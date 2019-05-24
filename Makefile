# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
WEB_BIN_NAME=albedo-web
CLI_BIN_NAME=albedo-cli
INSTALLER_BIN_NAME=albedo-installer

all: build

install: build
	cp -nv conf/config.yaml.example conf/config.yaml

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

release: release-cli release-web

release-cli: gotool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME)-linux -v cli.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME)-win.exe -v cli.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(CLI_BIN_NAME)-darwin -v cli.go

release-web: gotool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME)-linux -v web.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME)-win.exe -v web.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ./build/$(WEB_BIN_NAME)-darwin -v web.go

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	find ./build -type f | grep -v ".gitkeep" | xargs rm -v
	find . -type f -name .DS_Store -delete

gotool:
	gofmt -w .
