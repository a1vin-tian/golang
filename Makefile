export tag=v1.0

build:  $(shell find . -name '*.go')
	echo "building httpserver binary"
	mkdir -p bin/
	CGO_ENABLED=0 GOARCH=amd64 go build -o bin/ ./httpserver