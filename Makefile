test:
	go test ./internal/...

run:
	go run ./cmd/single/main.go

build:
	go build -o ./bin/single ./cmd/single/main.go

lint:
	golangci-lint run ./... --fast

help:
	@echo "You can contact with me on alameddinc@gmail.com"