.PHONY: run build test

run:
	go run ./...

build:
	go build -o bin/pz3-http ./cmd/server

test:
	go test ./... -v

clean:
	rm -rf bin/

fmt:
	go fmt ./...