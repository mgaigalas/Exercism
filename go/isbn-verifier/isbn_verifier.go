// Package isbn contains tools for validating ISBN codes.
package isbn

// IsValidISBN validates ISBN string.
func IsValidISBN(isbn string) bool {
	pos, sum := 10, 0

	for _, char := range isbn {
		switch {
		case char == '-':
			continue
		case char == 'X' && pos == 1:
			sum += 10
			pos--
		case char >= '0' && char <= '9':
			sum += (int(char) - '0') * pos
			pos--
		default:
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}
