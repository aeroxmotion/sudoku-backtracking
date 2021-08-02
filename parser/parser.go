package parser

import (
	"github.com/LosMaquios/sudoku-backtracking/sudoku"
	"math"
	"strconv"
	"strings"
)

// ParseSudokuInstructions
// parses the given row's instructions by interpreting all negative
// integers as a padding of `n` zeroes (internally managed as offset),
// e.g. given a row's instructions like this:
// "-4 5 -4"
// will generate the following row:
// "0 0 0 0 5 0 0 0 0"
func ParseSudokuInstructions(instructionRows [9]string) *sudoku.Sudoku {
	s := &sudoku.Sudoku{}

	for row, instructionRow := range instructionRows {
		commands := strings.Split(instructionRow, " ")
		col := 0

		for _, command := range commands {
			num, _ := strconv.Atoi(command)

			if num >= 0 {
				s.PutValue(row, col, num)
				col++
			} else {
				// Add repetitions as offset
				col += int(math.Abs(float64(num)))
			}
		}
	}

	return s
}
