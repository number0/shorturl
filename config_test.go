package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{}
	assert.Equal(cfg.FQDN, "")
}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{FQDN: "bar.com"}
	assert.Equal(cfg.FQDN, "bar.com")
}
