build:
	go-bindata -o server/assets.go -pkg server data/ && go build
test:
	go test ./...
up:
	make build && ./go-komonjo server
