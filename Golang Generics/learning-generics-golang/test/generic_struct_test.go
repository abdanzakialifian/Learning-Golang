package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[T]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Abdan",
		Second: "Alifian",
	}
	assert.Equal(t, "Hello Zaki", data.SayHello("Zaki"))
	assert.Equal(t, "Zaki", data.ChangeFirst("Zaki"))
}
