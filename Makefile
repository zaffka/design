.PHONY: build test lint

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X main.httpPort=$(HTTP_PORT)" -o main .

test:
	go test -v --race ./...

lint:
	golangci-lint run ./...