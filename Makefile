test:
	go test ./... -v

clean:
	rm -r bin/

build: test
	go build -ldflags="-s -w" -o bin/rooster ./cmd/rooster
	go build -ldflags="-s -w" -o bin/server ./cmd/server
