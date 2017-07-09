# shorturl
[![Build Status](https://travis-ci.org/prologic/shorturl.svg)](https://travis-ci.org/prologic/shorturl)
[![GoDoc](https://godoc.org/github.com/prologic/shorturl?status.svg)](https://godoc.org/github.com/prologic/shorturl)
[![Wiki](https://img.shields.io/badge/docs-wiki-blue.svg)](https://github.com/prologic/shorturl/wiki)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/shorturl)](https://goreportcard.com/report/github.com/prologic/shorturl)
[![Coverage](https://coveralls.io/repos/prologic/shorturl/badge.svg)](https://coveralls.io/r/prologic/shorturl)

shorturl is a web app that allows you to create short urls of much longer more
complex urls for easier sharing or embedding.

## Installation

### Source

```#!bash
$ go install github.com/prologic/shorturl/...
```

### OS X Homebrew

**Coming soon**

There is a formula provided that you can tap and install from
[prologic/homebrew-shorturl](https://github.com/prologic/homebrew-shorturl):

```#!bash
$ brew tap prologic/shorturl
$ brew install shorturl
```

**NB:** This installs the latest released binary; so if you want a more
recent unreleased version from master you'll have to clone the repository
and build yourself.

shorturl is still early days so contributions, ideas and expertise are
much appreciated and highly welcome!

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
