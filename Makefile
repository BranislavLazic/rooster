test:
	go test ./...

build: test
	go build -ldflags="-s -w" -o bin/rooster ./*.go
