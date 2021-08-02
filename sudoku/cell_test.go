package sudoku

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCell_HasValue(t *testing.T) {
	expectedValue := 1
	cell := Cell(expectedValue)

	for i := 0; i <= 9; i++ {
		assert.Equal(t, cell.HasValue(i), expectedValue == i)
	}
}

func TestCell_IsEmpty(t *testing.T) {
	var emptyCell Cell
	filledCell := Cell(2)

	assert.True(t, emptyCell.IsEmpty())
	assert.False(t, filledCell.IsEmpty())
}

func TestCell_String(t *testing.T) {
	var emptyCell Cell
	filledCell := Cell(1)

	assert.Equal(t, emptyCell.String(), "0")
	assert.Equal(t, filledCell.String(), "1")
}
