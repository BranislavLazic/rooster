APP=rooster
SERVER=rserver
VERSION := v0.1

.PHONY: clean bin test

all: clean test zip

clean:
	rm -rf bin release

test:
	go test ./...

zip: release/$(APP)_$(VERSION)_osx_x86_64.tar.gz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz release/$(APP)_$(VERSION)_linux_x86_32.tar.gz

bin: bin/osx_x86_64/$(APP) bin/linux_x86_64/$(APP) bin/osx_x86_32/$(APP) bin/linux_x86_32/$(APP)

release/$(APP)_$(VERSION)_osx_x86_64.tar.gz: bin/osx_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -C bin/osx_x86_64 $(APP)
	tar cfz release/$(SERVER)_$(VERSION)_osx_x86_64.tar.gz -C bin/osx_x86_64 $(SERVER)

bin/osx_x86_64/$(APP):
	GOOS=darwin GOARCH=amd64 go build -o bin/osx_x86_64/$(APP) ./cmd/rooster
	GOOS=darwin GOARCH=amd64 go build -o bin/osx_x86_64/$(SERVER) ./cmd/server

release/$(APP)_$(VERSION)_linux_x86_64.tar.gz: bin/linux_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -C bin/linux_x86_64 $(APP)
	tar cfz release/$(SERVER)_$(VERSION)_linux_x86_64.tar.gz -C bin/linux_x86_64 $(SERVER)

bin/linux_x86_64/$(APP):
	GOOS=linux GOARCH=amd64 go build -o bin/linux_x86_64/$(APP) ./cmd/rooster
	GOOS=linux GOARCH=amd64 go build -o bin/linux_x86_64/$(SERVER) ./cmd/server

release/$(APP)_$(VERSION)_osx_x86_32.tar.gz: bin/osx_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_32.tar.gz -C bin/osx_x86_32 $(APP)
	tar cfz release/$(SERVER)_$(VERSION)_osx_x86_32.tar.gz -C bin/osx_x86_32 $(SERVER)

bin/osx_x86_32/$(APP):
	GOOS=darwin GOARCH=386 go build -o bin/osx_x86_32/$(APP) ./cmd/rooster
	GOOS=darwin GOARCH=386 go build -o bin/osx_x86_32/$(SERVER) ./cmd/server

release/$(APP)_$(VERSION)_linux_x86_32.tar.gz: bin/linux_x86_32/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_32.tar.gz -C bin/linux_x86_32 $(APP)
	tar cfz release/$(SERVER)_$(VERSION)_linux_x86_32.tar.gz -C bin/linux_x86_32 $(SERVER)

bin/linux_x86_32/$(APP):
	GOOS=linux GOARCH=386 go build -o bin/linux_x86_32/$(APP) ./cmd/rooster
	GOOS=linux GOARCH=386 go build -o bin/linux_x86_32/$(SERVER) ./cmd/server