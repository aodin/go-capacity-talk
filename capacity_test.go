package capacity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	var Simple simple
	outputs := recurse([]int{}, reflect.TypeOf(Simple))
	assert.Len(t, outputs, 3)

	// A0
	assert.Equal(t, "A0", outputs[0].name)
	assert.Equal(t, []int{0}, outputs[0].index)

	// A1 - B0
	assert.Equal(t, "B0", outputs[1].name)
	assert.Equal(t, []int{1, 0}, outputs[1].index)

	// A1 - B1
	assert.Equal(t, "B1", outputs[2].name)
	assert.Equal(t, []int{1, 1}, outputs[2].index)
}

func TestComplex(t *testing.T) {
	var Nested nested
	outputs := recurse([]int{}, reflect.TypeOf(Nested))
	assert.Len(t, outputs, 3)

	// A0
	assert.Equal(t, "A0", outputs[0].name)
	assert.Equal(t, []int{0}, outputs[0].index)

	// A1 - B0 - C0 - D0
	assert.Equal(t, "D0", outputs[1].name)
	assert.Equal(t, []int{1, 0, 0, 0}, outputs[1].index) // Wat goes here

	// A1 - B0 - C0 - D1
	assert.Equal(t, "D1", outputs[2].name)
	assert.Equal(t, []int{1, 0, 0, 1}, outputs[2].index)
}

func TestCapacity(t *testing.T) {
	var slice []int
	assert.Len(t, slice, 0)
	assert.Equal(t, 0, cap(slice))

	slice = append(slice, 0, 1)
	assert.Len(t, slice, 2)
	assert.Equal(t, 2, cap(slice))

	slice = append(slice, 2)
	assert.Len(t, slice, 3)
	assert.Equal(t, 4, cap(slice))
}
