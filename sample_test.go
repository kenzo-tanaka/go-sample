package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	assert.Equal(t, add(1,1), 2)
}
