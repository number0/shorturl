FROM golang:alpine

EXPOSE 8000/tcp

ENTRYPOINT ["shorturl"]

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/shorturl
WORKDIR /go/src/shorturl

COPY . /go/src/shorturl

RUN go get -v -d
RUN go install -v
