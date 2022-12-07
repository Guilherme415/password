.PHONY: test

test:
	go test -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

coverage:
	@echo "Running project coverage..."
	@go test ./... -coverprofile fmtcoverage.html fmt
	@go test ./... -coverprofile cover.out
	@go tool cover -html=cover.out
	@go tool cover -html=cover.out -o cover.html
	@echo "Coverage completed successfully"

install:
	go mod vendor

vendor:
	go mod vendor

build:
	go build .

run:
	go run main.go

lint:
	go vet ./...
