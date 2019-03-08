test:
	go test ./... -v

build: test
	go build -ldflags="-s -w" -o bin/rooster ./*.go
