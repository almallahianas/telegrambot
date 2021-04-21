install:
	go get ./...

tidy:
	go mod tidy

test:
	go test -race ./...

build:
	go build -o dist/telegrambot cmd/main.go