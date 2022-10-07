package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	msg := hello()
	assert.Equal(t, msg, "Hello, World!", "Hello-World message is not equal")
}
