# shorturl
[![Build Status](https://travis-ci.org/prologic/shorturl.svg)](https://travis-ci.org/prologic/shorturl)
[![GoDoc](https://godoc.org/github.com/prologic/shorturl?status.svg)](https://godoc.org/github.com/prologic/shorturl)
[![Wiki](https://img.shields.io/badge/docs-wiki-blue.svg)](https://github.com/prologic/shorturl/wiki)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/shorturl)](https://goreportcard.com/report/github.com/prologic/shorturl)
[![Coverage](https://coveralls.io/repos/prologic/shorturl/badge.svg)](https://coveralls.io/r/prologic/shorturl)

shorturl is a web app that allows you to create smart bookmarks, commands and aliases by pointing your web browser's default search engine at a running instance. Similar to bunny1 or yubnub.

## Installation

### Source

```#!bash
$ go install github.com/prologic/shorturl/...
```

### OS X Homebrew

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
$ shorturl -bind 127.0.0.1:8000
```

Set your browser's default shorturl engine to http://localhost:8000/?q=%s

Then type `help` to view the main help page, `g foo bar` to perform a [Google](https://google.com) search for "foo bar" or `list` to list all available commands.

## License

MIT
