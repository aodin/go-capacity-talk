package capacity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// START OMIT
type deep struct {
	A0 int
	A1 struct {
		B0 struct {
			C0 struct {
				D0, D1 int
			}
		}
	}
}

func TestDeep(t *testing.T) {
	outputs := recurse([]int{}, reflect.TypeOf(deep{}))
	assert.Len(t, outputs, 3)

	assert.Equal(t, "A0", outputs[0].name)      // A0
	assert.Equal(t, []int{0}, outputs[0].index) // No Wat

	assert.Equal(t, "D0", outputs[1].name)               // A1 - B0 - C0 - D0
	assert.Equal(t, []int{1, 0, 0, 0}, outputs[1].index) // Wat goes here

	assert.Equal(t, "D1", outputs[2].name)               // A1 - B0 - C0 - D1
	assert.Equal(t, []int{1, 0, 0, 1}, outputs[2].index) // No wat
}

// END OMIT
