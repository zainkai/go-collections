GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

.PHONY: all test clean bench
all: test
test: 
	$(GOTEST) ./... -v
bench:
	$(GOTEST) -bench=. ./... -benchmem
deps:
	$(GOMOD) download