GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get -v
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

.PHONY: all

all: test

.PHONY: test
test:
	$(GOTEST) -v -covermode=count -coverprofile=coverage.out ./...

.PHONY: build
build: test
	$(GOBUILD)

.PHONY: install
install: test
	$(GOINSTALL)
