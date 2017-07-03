.PHONY: dev build clean

all: dev

dev: build
	./shorturl -bind 127.0.0.1:8000

build: clean
	go get ./...
	go build -o ./shorturl .

clean:
	rm -rf bin shorturl
