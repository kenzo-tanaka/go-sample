package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"./calc"
)

type Person struct {
	Lastname string
	Firstname string
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func TestAdd(t *testing.T) {
	assert.Equal(t, calc.Add(1,1), 2)
}

func TestPerson(t *testing.T) {
	person := new(Person)
	person.Lastname = "Kenzo"
	person.Firstname = "Tanaka"
	assert.Equal(t, person.Lastname, "Kenzo")
}
