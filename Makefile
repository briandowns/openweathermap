GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get -v
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

test:
	$(GOTEST) -v -cover ./...