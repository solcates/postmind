# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLINT=golint
BINARY_NAME=postmind
BINARY_UNIX=$(BINARY_NAME)
PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test build lint clean run deps

all: test build
build:
		CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -ldflags "-s -w"
lint:
		$(GOLINT) -set_exit_status $(PKGS)
vet:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOCMD) vet ./...
test:
		$(GOTEST) -short -race $(PKGS)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		go mod download


# Cross compilation
build-linux:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -ldflags  "-s -w"