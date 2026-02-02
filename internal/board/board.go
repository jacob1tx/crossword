package board

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

const width uint8 = 5
const height uint8 = 5

// Special Tiles
//
// The correct way to handle this would probably be enums.
// However, to put the letters and the square types in the same
// array, I would need to enumerate all of the letters, and that
// feels bad. Thus, I am "dumb enum-ing" these special tiles
// as constants.

const (
	MT uint8 = iota	// Empty
	DL 				// Double Letter
	DW 				// Double Word
	TL 				// Triple Letter
	TW 				// Triple Word
	ST 				// Starting Square
)

// board defines the game board
// It is initialized with theo
var board = [width][height]uint8{
	{MT, TW, MT, TW, MT},
	{TW, MT, MT, MT, TW},
	{MT, TW, ST, TW, MT},
	{TW, MT, MT, MT, TW},
	{MT, TW, MT, TW, MT},
}

func Set(x uint8, y uint8, letter uint8) (err error) {
	errs := validateArrayInputs(x, y)
	if errs != nil {
		return errs
	}

	board[x][y] = letter
	
	log.Info().Msgf("[%d,%d] = %c", x, y, letter)
	log.Debug().Msgf("%v", board)

	return nil
}

func Get(x uint8, y uint8) (letter uint8, err error) {
	errs := validateArrayInputs(x, y)
	if errs != nil {
		return MT, errs
	}
	return board[x][y], nil
}

func validateArrayInputs(x uint8, y uint8) (error) {
	var errs error

	if x >= width {
		errs = errors.Join(fmt.Errorf("x, %d, must satisfy x < width (%d).", x, width), errs)
	}
	if y >= height {
		errs = errors.Join(fmt.Errorf("y, %d, must satisfy y < height (%d).", y, height), errs)
	}

	return errs
}
