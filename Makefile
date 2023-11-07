export AOC_SESSION ?= $(shell cat .aoc-session)

run:
	@go run main.go run $(key)

all:
	@go run main.go all

stubs:
	@go run main.go stubs $(key)

list:
	@go run main.go list

test: gotestsum
	gotestsum --format=testdox -- -coverprofile=coverage.out ./...

gotestsum:
	go install gotest.tools/gotestsum@latest

lint:
	golangci-lint run ./...
