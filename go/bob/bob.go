// Package bob contains tools for matching remark patterns to responses.
package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	Remark string
}

// Hey contains tools matching remarks as patterns to responses.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	isAllCaps := isUpperCase(remark)
	isQuestion := strings.HasSuffix(remark, "?")

	switch {
	case isAllCaps && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isAllCaps && !isQuestion:
		return "Whoa, chill out!"
	case !isAllCaps && isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

func isUpperCase(str string) bool {
	var isUpperCase bool
	for _, r := range str {
		if !unicode.IsLetter(r) {
			continue
		}

		if unicode.IsLower(r) {
			isUpperCase = false
			break
		}
		isUpperCase = true
	}
	return isUpperCase
}
