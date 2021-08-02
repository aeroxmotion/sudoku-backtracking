package main

import (
	"github.com/LosMaquios/sudoku-backtracking/parser"
	"github.com/LosMaquios/sudoku-backtracking/solver"
)

func main() {
	s := parser.ParseSudokuInstructions([9]string{
		"0 3 2 -3 8",
		"0 6 -4 1",
		"5 -7 7",

		"0 1 -3 4 0 9 0",
		"-2 4 -2 6",
		"0 9 -2 1 0 3",

		"7 -4 5",
		"8 5 0 9 4 -3 6",
		"-4 8 -2 2",
	})

	println("===== Parsed sudoku =====\n" + s.String())
	println()

	solver.SolveSudoku(s)

	println("===== Solved sudoku =====\n" + s.String())
}
