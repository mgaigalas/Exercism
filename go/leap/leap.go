// Package leap contains tools for calculating if specific year is a leap year
package leap

// IsLeapYear calculates if specific year is a leap year
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
