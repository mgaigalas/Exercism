// Package clock contains tools for applying arithmetic operation on custom Clock implementation.
package clock

import "fmt"

type Clock int

// New creates new Clock.
func New(h, m int) Clock {
	totalMinutes := (60*(h%24) + m) % (24 * 60)
	if totalMinutes < 0 {
		totalMinutes = 24*60 + totalMinutes
	}
	return Clock(totalMinutes)
}

// Add adds values to Clock.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract subtracts values from Clock.
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

// String represents time defined as a Clock.
func (c Clock) String() string {
	hours := (int(c) / 60) % 24
	minutes := int(c) % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
