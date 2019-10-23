GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY=go-collections
BINARY_TEST=$(BINARY)-test

.PHONY: all test clean bench
all: test
test: 
	$(GOTEST) ./... -v
bench:
	$(GOTEST) -bench=. ./... -benchmem
deps:
	$(GOMOD) download
build-test:
	$(GOTEST) -c ./... -o $(BINARY_TEST)
clean:
	rm $(BINARY_TEST)
