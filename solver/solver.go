package solver

import "github.com/LosMaquios/sudoku-backtracking/sudoku"

func SolveSudoku(s *sudoku.Sudoku) {
	validSudoku := SolveSudokuRecursive(s, 0, 0)

	if !validSudoku {
		println("[DEBUG]: Invalid sudoku!!!")
	}
}

func SolveSudokuRecursive(s *sudoku.Sudoku, rowIndex, colIndex int) bool {
	if colIndex == 9 {
		colIndex = 0
		rowIndex++
	}

	if rowIndex >= 9 {
		return true
	}

	if !s.IsEmpty(rowIndex, colIndex) {
		return SolveSudokuRecursive(s, rowIndex, colIndex + 1)
	}

	for n := 1; n <= 9; n++ {
		if s.IsValid(rowIndex, colIndex, n) {
			s.PutValue(rowIndex, colIndex, n)

			if SolveSudokuRecursive(s, rowIndex, colIndex + 1) {
				return true
			}

			s.PutValue(rowIndex, colIndex, 0)
		}
	}

	return false
}
