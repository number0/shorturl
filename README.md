# shorturl
[![GoDoc](https://godoc.org/github.com/prologic/shorturl?status.svg)](https://godoc.org/github.com/prologic/shorturl)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/shorturl)](https://goreportcard.com/report/github.com/prologic/shorturl)

shorturl is a web app that allows you to create short urls of much longer more
complex urls for easier sharing or embedding.

## Installation

### Source

```#!bash
$ go install github.com/prologic/shorturl/...
```

## Usage

Run shorturl:

```#!bash
$ shorturl
```

Then visit: http://localhost:8000/

## Configuration

By default shorturl stores urls in `urls.db` in the local directory. This can
be configured with the `-dbpath /path/to/urls.db` option.

shorturl also displays an absolute url after creating and uses the value of
`-baseurl` (*default: `""`*) for display purposes. This is useful for copying
and pasting the shorturl.

## License

MIT
