// Package etl contains tools for transforming points-to-letters one-to-many relationships to one-to-one relationships.
package etl

import "strings"

// Transform transforms one-to-many into one-to-one score-to-letters data structure
func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for key, val := range in {
		for _, char := range val {
			res[strings.ToLower(char)] = key
		}
	}
	return res
}
