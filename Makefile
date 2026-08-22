.DEFAULT_GOAL := build

#MAKEFLAGS += --silent

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build: vet
	go build -o c64char ./cmd/c64char.go

test: build
	go test -v -count=1 ./...

run: build
	./c64char.go
