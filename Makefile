proj=me_exporter

TARGETDIR = ./deploy
GO_OBJS   = $(wildcard **/*.go)

ifeq ($(OS),Windows_NT)
else
    UNAME_S := $(shell uname -s)
    UNAME_P := $(shell uname -p)
    date := $(shell date +%Y%m%d%H%M%S)
endif

sha1ver := $(shell git rev-parse HEAD)
gittag  := $(shell git describe --tag)

LDFLAGS=-ldflags "-X main.sha1ver=$(sha1ver) -X main.Version=$(gittag)"

.PHONY: all build-linux-arm64 build-linux-x86_64 build-macos-arm64 build-macos-x86_64

build-linux-arm64: $(GO_OBJS)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(TARGETDIR)/$(proj)-linux-arm64

build-linux-x86_64: $(GO_OBJS)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(TARGETDIR)/$(proj)-linux-x86_64

build-macos-arm64: $(GO_OBJS)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(TARGETDIR)/$(proj)-macos-arm64

build-macos-x86_64: $(GO_OBJS)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(TARGETDIR)/$(proj)-macos-x86_64

all: vet test build

vet:
	go vet -all ./

test:
	go.exe test -timeout 30s ./...

build: build-linux-arm64 build-linux-x86_64 build-macos-arm64 build-macos-x86_64
