package main

import "errors"

const (
	rows    = 9
	columns = 9
)

var (
	ErrBounds  = errors.New("out of bounds")
	ErrDigit   = errors.New("invalid digit")
	ErrSameRow = errors.New("Same Row")
	ErrSameCol = errors.New("Same Col")
)

type Grid [rows][columns]int8

func NewSudoku(g [rows][columns]int8) Grid {
	return Grid(g)
}

func (g *Grid) Clear(row, column int) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	g[row][column] = 0
	return nil
}

// Set a digit on a Sudoku grid
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return ErrDigit
	}
	if samerow(g, digit, row) {
		return ErrSameRow
	}
	if samecolumn(g, digit, column) {
		return ErrSameCol
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func samerow(g *Grid, digit int8, row int) bool {
	for _, v := range g[row] {
		if digit == v {
			return true
		}
	}
	return false
}

func samecolumn(g *Grid, digit int8, column int) bool {
	for i := 0; i < rows; i++ {
		if g[i][column] == digit {
			return true
		}
	}
	return false
}

func samesub(row, column int, g *Grid, digit int8) {
	sub_row := row % 3
	sub_column := column % 3
	for i := 0; i < sub_row; i++ {

	}
}

func main() {
	s := NewSudoku(Grid)
}
