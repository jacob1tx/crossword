package game

import (
	"errors"
	"fmt"
)
		
const width int = 5
const height int = 5

var board [width][height]rune

func set(x int, y int, letter rune) (err error) {
	errs := validateArrayInputs(x, y)
	if errs != nil {
		return errs
	}
	board[x][y] = letter
	return nil
}

func get(x int, y int) (letter rune, err error) {
	errs := validateArrayInputs(x, y)
	if errs != nil {
		return -1, errs
	}
	return board[x][y], nil
}

func validateArrayInputs(x int, y int) (error) {
	var errs error

	if x >= width || x < 0 {
		errs = errors.Join(fmt.Errorf("x, %d, must satisfy 0 <= x < width (%d).", x, width), errs)
	}
	if y >= height || y < 0 {
		errs = errors.Join(fmt.Errorf("y, %d, must satisfy 0 <= y < height (%d).", y, height), errs)
	}

	return errs
}
