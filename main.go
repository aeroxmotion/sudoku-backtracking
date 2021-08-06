package main

import (
	"github.com/LosMaquios/sudoku-backtracking/parser"
	"github.com/LosMaquios/sudoku-backtracking/solver"
)

// Test by showing solved sudoku
func main() {
	s := parser.ParseSudokuInstructions([9]string{
		"8",
		"-2 3 6",
		"0 7 -2 9 0 2",

		"0 5 -3 7",
		"-4 4 5 7",
		"-3 1 -3 3",

		"-2 1 -4 6 8",
		"-2 8 5 -3 1",
		"0 9 -4 4",
	})

	solver.SolveSudoku(s)

	println("\n> Solved sudoku <\n" + s.String())
}
