export AOC_SESSION ?= $(shell cat .aoc-session)

run:
	@go run main.go run $(key)

stubs:
	@go run main.go stubs $(key)

list:
	@go run main.go list

test: gotestsum
	gotestsum --format=testname -- -coverprofile=coverage.out ./...

gotestsum:
	go install gotest.tools/gotestsum@latest
