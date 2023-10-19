// Package isogram contains tools for calculating if word is an isogram
package isogram

import "unicode"

// IsIsogram checks if word is an isogram
func IsIsogram(word string) bool {
	dictionary := make(map[rune]int, len(word))
	for _, val := range word {
		if !unicode.IsLetter(val) {
			continue
		}
		char := unicode.ToLower(val)
		if dictionary[char] > 0 {
			return false
		}
		dictionary[char]++
	}
	return true
}
