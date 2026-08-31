.DEFAULT_GOAL := build

#MAKEFLAGS += --silent

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build: vet
	go build ./

test: build
	go test -v -count=1 ./...
