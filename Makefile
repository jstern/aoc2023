export AOC_SESSION ?= $(shell cat .aoc-session)

run:
	@go run main.go run $(key)

submit:
	@go run main.go submit $(key)

all:
	@go run main.go all

stubs:
	@go run main.go stubs $(key)

list:
	@go run main.go list


gts := $(shell which gotestsum)
test: $(gts)
	$(gts) --format=testdox -- -coverprofile=coverage.out ./...

$(gts):
	go install gotest.tools/gotestsum@latest

lint: fmt
	golangci-lint run ./...

fmt:
	go fmt ./...

cov:
	go tool cover -html=coverage.out
