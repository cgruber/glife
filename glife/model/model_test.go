package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewArrayBitMatrixFromCopy tests the creation of an ArrayBitMatrix from the data of a BitMatrix
func TestNewArrayBitMatrixFromCopy(t *testing.T) {
	initial := NewArrayBitMatrix(5, 5)
	initial.Set(2, 3, true)
	initial.Set(4, 4, true)
	assert.True(t, initial.Get(2, 3), "Value should be changed to true")
	assert.True(t, initial.Get(4, 4), "Value should be changed to true")
	dupe := NewArrayBitMatrixFromExisting(initial)
	assert.NotSame(t, initial, dupe, "Initial and copied matrix should not have reference-equality")
	assert.Equal(t, initial, dupe, "Initial and copied matrix should have value-equality")
}

// TestNewArrayBitMatrix tests the creation of an empty ArrayBitMatrix
func TestNewArrayBitMatrix(t *testing.T) {
	actual := NewArrayBitMatrix(12, 9)
	assert.NotNil(t, actual, "ArrayBitMatrix should not be null")
	for x := uint64(0); x < 12; x++ {
		for y := uint64(0); y < 9; y++ {
			assert.False(t, actual.Get(x, y), "All values should be initialized to false.")
		}
	}
}

// TestGetAndSetOnArrayBitMatrix tests the creation of an empty ArrayBitMatrix
func TestGetAndSetOnArrayBitMatrix(t *testing.T) {
	abm := NewArrayBitMatrix(12, 9)
	assert.NotNil(t, abm, "ArrayBitMatrix should not be null")
	assert.False(t, abm.Get(2, 3), "Value should be initialized to false.")
	abm.Set(2, 3, true)
	assert.True(t, abm.Get(2, 3), "Value should be changed to true")
}
