package logs

import "unicode/utf8"

var unicodeRuneDictionary = map[string]string{
	"\u2757":     "recommendation",
	"\U0001f50d": "search",
	"\u2600":     "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if val, exists := unicodeRuneDictionary[string(char)]; exists {
			return val
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var fixedRunes []rune
	for _, char := range log {
		if char == oldRune {
			fixedRunes = append(fixedRunes, newRune)
			continue
		}
		fixedRunes = append(fixedRunes, char)
	}
	return string(fixedRunes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
