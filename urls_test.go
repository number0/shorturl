package main

import (
	"testing"

	"github.com/boltdb/bolt"

	"github.com/stretchr/testify/assert"
)

func TestZeroURL(t *testing.T) {
	assert := assert.New(t)

	u := URL{}
	assert.Equal(u.ID(), "")
	assert.Equal(u.URL(), "")
}

func TestURLSaveLookup(t *testing.T) {
	assert := assert.New(t)

	db, _ = bolt.Open("test.db", 0600, nil)
	defer db.Close()

	URL{id: "asdf", url: "https://localhost"}.Save()

	u, ok := LookupURL("asdf")
	assert.True(ok)
	assert.Equal(u.id, "asdf")
	assert.Equal(u.url, "https://localhost")
}
