// Package luhn is a small library used for validating string (i.e. credit card numbers) against Luhs algorithm
package luhn

import (
	"strings"
	"unicode"
)

// Valid validates string against Luhs algorithm
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	var sum int
	parity := len(id) % 2
	for i := 0; i < len(id); i++ {
		char := rune(id[i])
		if !unicode.IsNumber(char) {
			return false
		}

		num := int(char - '0')
		switch {
		case i%2 != parity:
			sum += num
		case num > 4:
			sum += num*2 - 9
		default:
			sum += num * 2
		}
	}
	return sum%10 == 0
}
