// Package pangram is a small library for checking if string is considered to be a pangram
package pangram

import (
	"strings"
)

// IsPangram checks if string is a pangram string
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(input, char) {
			return false
		}
	}
	return true
}
