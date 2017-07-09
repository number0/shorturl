package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZeroConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{}
	assert.Equal(cfg.baseURL, "")
}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{baseURL: "http://localhost:8000/"}
	assert.Equal(cfg.baseURL, "http://localhost:8000/")
}
