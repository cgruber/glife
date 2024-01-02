package model

import (
	"fmt"
	"testing"
)

// TestNewArrayBitMatrixFromCopy tests the creation of an ArrayBitMatrix from the data of a BitMatrix
func TestNewArrayBitMatrixFromCopy(t *testing.T) {
	fmt.Println("Yay, a test!")
}

// TestNewArrayBitMatrix tests the creation of an empty ArrayBitMatrix
func TestNewArrayBitMatrix(t *testing.T) {
	actual := NewArrayBitMatrix(12, 9)
	if actual == nil {
		t.Errorf("ArrayBitMatrix should not be null")
	}
	fmt.Println(actual)
	for x := uint64(0); x < 12; x++ {
		for y := uint64(0); y < 9; y++ {
			if actual.Get(x, y) != false {
				t.Errorf("All values should be initialized to zero.")
			}
		}
	}
}
