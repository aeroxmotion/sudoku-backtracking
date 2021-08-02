package main

import (
	"github.com/LosMaquios/sudoku-backtracking/parser"
	"github.com/LosMaquios/sudoku-backtracking/solver"
)

func main() {
	s := parser.ParseSudokuInstructions([9]string{
		"-2 2 0 8 -2 7 9",
		"0 9 -4 1 -2",
		"7 0 4 -2 9 -3",

		"9 1 -7",
		"-5 8 -2 3",
		"0 8 -3 7 2 5 0",

		"-4 2 -2 4 5",
		"-3 4 -3 9 0",
		"3 4 -6 6",
	})

	println("===== Parsed sudoku =====\n" + s.String())
	println()

	solver.SolveSudoku(s)

	println("===== Solved sudoku =====\n" + s.String())
}
