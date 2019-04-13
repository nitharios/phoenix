USER=nitharios
BINARY=main
VERSION=0.0.1

BIN_PATH=bin

DIR=$(shell pwd)
RUN_ENV=$(shell cat .env | xargs)

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get -u -v -d
GOTEST=$(GOCMD) test

GOARCH=amd64
OS=linux

all: build

build:
	$(GOBUILD) -o $(BIN_PATH)/$(BINARY) -v

build-amd64:
	env CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(GOARCH) $(GOBUILD) -a -o $(BIN_PATH)/$(BINARY) -v .

env:
	cp ./.env.sample .env