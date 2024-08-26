tests:
	go test -v ./... -race
.PHONY: tests


build:
	go build -o bin/ ./...