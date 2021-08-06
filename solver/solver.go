package solver

import "github.com/LosMaquios/sudoku-backtracking/sudoku"

func SolveSudoku(s *sudoku.Sudoku) {
	validSudoku := solveSudokuRecursive(s, 0, 0)

	if !validSudoku {
		println("[DEBUG]: Invalid sudoku!!!")
	}
}

// y = row's index
// x = column's index
func solveSudokuRecursive(sudoku *sudoku.Sudoku, y, x int) bool {
	if x == 9 {
		x = 0
		y++
	}

	if y == 9 {
		return true
	}

	if !sudoku.IsEmpty(y, x) {
		return solveSudokuRecursive(sudoku, y, x + 1)
	}

	for n := 1; n <= 9; n++ {
		if sudoku.IsValid(y, x, n) {
			sudoku.PutValue(y, x, n)

			if solveSudokuRecursive(sudoku, y, x + 1) {
				return true
			}

			sudoku.PutValue(y, x, 0)
		}
	}

	return false
}
