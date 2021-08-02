package parser

import (
	"github.com/LosMaquios/sudoku-backtracking/sudoku"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSudokuInstructions(t *testing.T) {
	parsedSudoku := ParseSudokuInstructions([9]string{
		"-2 5 -2 8 3 9",
		"0 3",
		"-3 7 -3 8",

		"-2 4 5 -2 6 0 2",
		"6 1",
		"2 -3 4 9",

		"-5 2 4 0 5",
		"-2 9 0 8",
		"5 6",
	})

	expected := sudoku.Sudoku{
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

	assert.ElementsMatch(t, *parsedSudoku, expected)
}
