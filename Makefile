.PHONY: all pre_build lint build 

all:
	$(MAKE) pre_build 
	$(MAKE) lint
	$(MAKE) build 

pre_build:
	go generate ./...
	go mod tidy

lint:
	gofmt -s -w ./
	goimports -w ./

build:
	go build

