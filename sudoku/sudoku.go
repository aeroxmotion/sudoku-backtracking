package sudoku

import (
	"math"
	"strings"
)

type Sudoku [9][9]Cell

func New() *Sudoku {
	return &Sudoku{}
}

func (s Sudoku) IsValid(rowIndex, colIndex, value int) bool {
	return !s.InSquare(rowIndex, colIndex, value) &&
		!s.InCross(rowIndex, colIndex, value)
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

func (s Sudoku) InCross(rowIndex, colIndex, value int) bool {
	for i, row := range s {
		if
			// In row?
			row[colIndex].HasValue(value) ||
			// In column?
			s[rowIndex][i].HasValue(value) {
			return true
		}
	}

	return false
}

func (s Sudoku) IsEmpty(rowIndex, colIndex int) bool {
	return s[rowIndex][colIndex].IsEmpty()
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
