// Package rotationalcipher contains tools for encrypting text with rotational cypher.
package rotationalcipher

import (
	"strings"
)

// RotationalCipher encrypts string with rotational cypher with provided key.
func RotationalCipher(plain string, shiftKey int) string {
	var sb strings.Builder
	for _, char := range plain {
		sb.WriteRune(rotate(char, shiftKey))
	}
	return sb.String()
}

func rotate(char rune, key int) rune {
	switch {
	case char >= 'a' && char <= 'z':
		return rune((int(char-'a')+key)%26) + 'a'
	case char >= 'A' && char <= 'Z':
		return rune((int(char-'A')+key)%26) + 'A'
	default:
		return char
	}
}
