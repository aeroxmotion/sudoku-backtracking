package solver

import (
	"github.com/LosMaquios/sudoku-backtracking/sudoku"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	s := &sudoku.Sudoku{
		{0, 0, 5,   0, 0, 8,   3, 9, 0},
		{0, 3, 0,   0, 0, 0,   0, 0, 0},
		{0, 0, 0,   7, 0, 0,   0, 8, 0},

		{0, 0, 4,   5, 0, 0,   6, 0, 2},
		{6, 1, 0,   0, 0, 0,   0, 0, 0},
		{2, 0, 0,   0, 4, 9,   0, 0, 0},

		{0, 0, 0,   0, 0, 2,   4, 0, 5},
		{0, 0, 9,   0, 8, 0,   0, 0, 0},
		{5, 6, 0,   0, 0, 0,   0, 0, 0},
	}

	solvedSudoku := sudoku.Sudoku{
		{7, 4, 5,   2, 1, 8,   3, 9, 6},
		{8, 3, 2,   4, 9, 6,   1, 5, 7},
		{1, 9, 6,   7, 5, 3,   2, 8, 4},

		{9, 8, 4,   5, 3, 1,   6, 7, 2},
		{6, 1, 3,   8, 2, 7,   5, 4, 9},
		{2, 5, 7,   6, 4, 9,   8, 3, 1},

		{3, 7, 8,   9, 6, 2,   4, 1, 5},
		{4, 2, 9,   1, 8, 5,   7, 6, 3},
		{5, 6, 1,   3, 7, 4,   9, 2, 8},
	}

	SolveSudoku(s)

	assert.ElementsMatch(t, solvedSudoku, *s)
}
