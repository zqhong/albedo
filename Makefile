WEB_BIN_NAME=albedo-web
CLI_BIN_NAME=albedo-cli
PROJECT_NAME=github.com/zqhong/albedo

VERSION=$(shell git describe --tags || echo "unkown version")
BUILDTIME=$(shell date -u)
GOBUILD=go build -ldflags '-X "github.com/zqhong/albedo/constant.Version=$(VERSION)" \
		-X "github.com/zqhong/albedo/constant.BuildTime=$(BUILDTIME)" \
		-w -s'

all: build

install: build
	cp -nv conf/config.yaml.example conf/config.yaml
	find . -name "*.go" | xargs -I {} sed -i "" "s#github.com/zqhong/albedo/#$(PROJECT_NAME)/#g" {}

build: build-web build-cli

build-web: gotool
	$(GOBUILD) -v -o ./build/$(WEB_BIN_NAME) -tags=jsoniter web.go

build-cli: gotool
	$(GOBUILD) -v -o ./build/$(CLI_BIN_NAME) -tags=jsoniter cli.go

run-web: build-web
	./build/$(WEB_BIN_NAME)

run-cli: build-cli
	./build/$(CLI_BIN_NAME)

release: release-cli release-web

release-cli: gotool
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(CLI_BIN_NAME)-linux -v cli.go
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./build/$(CLI_BIN_NAME)-win.exe -v cli.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./build/$(CLI_BIN_NAME)-darwin -v cli.go

release-web: gotool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(WEB_BIN_NAME)-linux -v web.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./build/$(WEB_BIN_NAME)-win.exe -v web.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./build/$(WEB_BIN_NAME)-darwin -v web.go

test:
	go test -v ./

clean:
	go clean
	find ./build -type f | grep -v ".gitkeep" | xargs rm -v
	find . -type f -name .DS_Store -delete

gotool:
	gofmt -w .
