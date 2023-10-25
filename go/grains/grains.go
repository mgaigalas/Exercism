// Package grains is a library containing tools for calculating grains on particular squares.
package grains

import (
	"errors"
)

// Square calculates number of grains on a square.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square")
	}
	return 1 << (number - 1), nil
}

// Total calculates number of grains on 64th square
func Total() uint64 {
	return 1<<64 - 1
}
