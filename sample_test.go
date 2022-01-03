package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"./calc"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	assert.Equal(t, calc.Add(1,1), 2)
}
