.PHONY: dev build clean

all: dev

dev: build
	./shorturl

build: clean
	go get ./...
	go build .

test:
	go test ./...

clean:
	rm -rf shorturl
