GO = go

.PHONY: all

all: test

.PHONY: test
test:
	$(GO) test -v -covermode=count -coverprofile=coverage.out ./...

