test:
	go test ./internal/...

run:
	go run ./cmd/singleservice/main.go

build:
	go build -o ./bin/single ./cmd/single/main.go

lint:
	golangci-lint run --config ./.golangci.yml

help:
	@echo "You can contact with me on alameddinc@gmail.com"