package sudoku

import (
	"math"
	"strings"
)

// Sudoku borrowed from:
// https://medium.com/swlh/backtracking-algorithm-to-solve-sudoku-puzzle-in-javascript-732aedcf5e2
type Sudoku [9][9]Cell

func (s Sudoku) IsValid(rowIndex, colIndex, value int) bool {
	return !s.InSquare(rowIndex, colIndex, value) &&
		!s.InRow(rowIndex, value) &&
		!s.InColumn(colIndex, value)
}

func (s Sudoku) InSquare(rowIndex, colIndex, value int) bool {
	boxRow, boxCol := getLineBox(rowIndex), getLineBox(colIndex)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if s[boxRow + r][boxCol + c].HasValue(value) {
				return true
			}
		}
	}

	return false
}

func (s Sudoku) InRow(rowIndex, value int) bool {
	for _, cell := range s[rowIndex] {
		if cell.HasValue(value) {
			return true
		}
	}

	return false
}

func (s Sudoku) InColumn(colIndex, value int) bool {
	for _, row := range s {
		if row[colIndex].HasValue(value) {
			return true
		}
	}

	return false
}

func (s *Sudoku) PutValue(rowIndex, colIndex, value int) {
	s[rowIndex][colIndex] = Cell(value)
}

func (s Sudoku) String() string {
	var str string

	for r, row := range s {
		if r != 0 && r % 3 == 0 {
			str += strings.Repeat("-", 9 + 2 /* <- padding */) + "\n"
		}

		for c, cell := range row {
			if c != 0 && c % 3 == 0 {
				str += "|"
			}

			str += cell.String()
		}

		str += "\n"
	}

	return str
}

func getLineBox(index int) int {
	return int(math.Floor(float64(index / 3)) * 3)
}
