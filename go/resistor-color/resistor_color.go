// Package resistorcolor contains utilities for decoding resistor color codes
package resistorcolor

import "strings"

// Colors should return the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for idx, val := range Colors() {
		if strings.ToLower(val) == color {
			return idx
		}
	}
	return -1
}
