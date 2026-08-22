.DEFAULT_GOAL := build

#MAKEFLAGS += --silent

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build: vet
	go build -v -o c64char ./cmd/c64char.go

test: build
	go test -v ./...

run: build
	./c64char.go
