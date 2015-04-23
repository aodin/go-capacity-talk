package capacity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
