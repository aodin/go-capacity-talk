package capacity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// START OMIT
type shallow struct {
	A0 int
	A1 struct {
		B0, B1 int
	}
}

func TestShallow(t *testing.T) {
	outputs := recurse([]int{}, reflect.TypeOf(shallow{}))
	assert.Len(t, outputs, 3)

	assert.Equal(t, "A0", outputs[0].name) // A0
	assert.Equal(t, []int{0}, outputs[0].index)

	assert.Equal(t, "B0", outputs[1].name) // A1 - B0
	assert.Equal(t, []int{1, 0}, outputs[1].index)

	assert.Equal(t, "B1", outputs[2].name) // A1 - B1
	assert.Equal(t, []int{1, 1}, outputs[2].index)
}

// END OMIT
