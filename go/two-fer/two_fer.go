// Package twofer provides tools for
// generating "two fer" string generation.
package twofer

import "fmt"

// ShareWith generates "two fer" string based on provided name
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
