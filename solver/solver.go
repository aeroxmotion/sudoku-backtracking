package solver

import "github.com/LosMaquios/sudoku-backtracking/sudoku"

func SolveSudoku(s *sudoku.Sudoku) {
	row, col := getNextEmptySpot(s)

	if row == -1 {
		return
	}

	for n := 1; n <= 9; n++ {
		if s.IsValid(row, col, n) {
			s.PutValue(row, col, n)
			SolveSudoku(s)
		}
	}

	if nextRow, _ := getNextEmptySpot(s); nextRow != -1 {
		s.PutValue(row, col, 0)
	}
}

func getNextEmptySpot(s *sudoku.Sudoku) (int, int) {
	for r, row := range s {
		for c, cell := range row {
			if cell.IsEmpty() {
				return r, c
			}
		}
	}

	return -1, -1
}
