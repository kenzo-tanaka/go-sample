package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct{ Name string }

func (p Person) Greet() {
	return "Foo"
}

func TestGreet(t *testing.T) {
	p := Person{Name: "Taro"}
	assert.Equal(t, p.Greet(), "Foo")
}
