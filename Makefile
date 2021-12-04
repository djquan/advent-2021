build:
	go build ./...

test:
	go test ./...

lint:
	golangci-lint run