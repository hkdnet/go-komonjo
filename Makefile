build:
	go build
test:
	go test ./...
up:
	go build && ./go-komonjo server
