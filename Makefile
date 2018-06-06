# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=buildbot

all: test build
build:
    $(GOBUILD) -o $(BINARY_NAME) -v ./src/...
test:
    $(GOTEST) -v ./src/...
clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)
run:
    $(GOBUILD) -o $(BINARY_NAME) -v ./src/...
    ./$(BINARY_NAME)
deps:
    dep ensure -update
setup:
    $(GOGET) -u github.com/golang/dep/cmd/dep
