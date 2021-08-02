package sudoku

import "strconv"

type Cell int

func (c Cell) IsEmpty() bool {
	return int(c) == 0
}

func (c Cell) HasValue(otherValue int) bool {
	return c == Cell(otherValue)
}

func (c Cell) String() string {
	return strconv.Itoa(int(c))
}
