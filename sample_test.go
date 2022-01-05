package main

import (
	"testing"

	"./calc"
	"github.com/stretchr/testify/assert"
)

// 関数の外では var が必要
// k := 3

const Pi = 3.14

type Person struct {
	Lastname  string
	Firstname string
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func defaultValue() int {
	var i int
	return i
}

func convert(num int) float64 {
	return float64(num)
}

func TestAdd(t *testing.T) {
	assert.Equal(t, calc.Add(1, 1), 2)
}

func TestPerson(t *testing.T) {
	person := new(Person)
	person.Lastname = "Kenzo"
	person.Firstname = "Tanaka"
	assert.Equal(t, person.Lastname, "Kenzo")
}

func TestDefaultValue(t *testing.T) {
	assert.Equal(t, defaultValue(), 0)
}

func TestConvert(t *testing.T) {
	assert.Equal(t, convert(1), 1.0)
}
