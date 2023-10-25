// Package darts is a library for calculating scores in darts game
package darts

import "math"

// Score returns dart score based on provided coordinates
func Score(x, y float64) int {
	distance := math.Sqrt((0-x)*(0-x) + (0-y)*(0-y))
	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}
