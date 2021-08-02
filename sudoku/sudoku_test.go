package sudoku

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var s = &Sudoku{
	{1, 0, 0,   0, 0, 0,   0, 0, 0},
	{0, 2, 0,   0, 0, 0,   0, 0, 0},
	{0, 0, 3,   0, 0, 0,   0, 0, 0},

	{0, 0, 0,   4, 0, 0,   0, 0, 0},
	{0, 0, 0,   0, 5, 0,   0, 0, 0},
	{0, 0, 0,   0, 0, 6,   0, 0, 0},

	{0, 0, 0,   0, 0, 0,   7, 0, 0},
	{0, 0, 0,   0, 0, 0,   0, 8, 0},
	{0, 0, 0,   0, 0, 0,   0, 0, 9},
}

func TestSudoku_IsEmpty(t *testing.T) {
	assert.True(t, s.IsEmpty(0, 1))
	assert.False(t, s.IsEmpty(7, 7))
}

func TestSudoku_InRow(t *testing.T) {
	assert.True(t, s.InRow(0, 1))
	assert.False(t, s.InRow(2, 2))
}

func TestSudoku_InColumn(t *testing.T) {
	assert.True(t, s.InColumn(4, 5))
	assert.False(t, s.InColumn(7, 3))
}

func TestSudoku_InSquare(t *testing.T) {
	assert.True(t, s.InSquare(0, 0, 3))
	assert.False(t, s.InSquare(3, 3, 3))
}

func TestSudoku_IsValid(t *testing.T) {
	assert.True(t, s.IsValid(0, 0, 4))
	assert.True(t, s.IsValid(6, 6, 4))
}

func TestSudoku_PutValue(t *testing.T) {
	sCopy := *s

	assert.True(t, sCopy.IsEmpty(0, 1))

	sCopy.PutValue(0, 1, 9)

	assert.Equal(t, 9, int(sCopy[0][1]))

	// Assert original sudoku's object have not been modified
	assert.True(t, s.IsEmpty(0, 1))
}

func TestSudoku_String(t *testing.T) {
	expectedRows := [11]string{
		"100|000|000",
		"020|000|000",
		"003|000|000",
		"-----------",
		"000|400|000",
		"000|050|000",
		"000|006|000",
		"-----------",
		"000|000|700",
		"000|000|080",
		"000|000|009",
	}

	// TODO: We should trim trailing line break?
	expectedFormat := strings.Join(expectedRows[:], "\n") + "\n"

	assert.Equal(t, expectedFormat, s.String())
}
