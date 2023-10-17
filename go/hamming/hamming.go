// Package hamming calculates hamming distance for provided DNA strains
package hamming

import "errors"

// Distance calculates hamming distance for provided DNA strains
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strain lengths mismatch")
	}

	var count int
	for idx := range a {
		if a[idx] != b[idx] {
			count++
		}
	}
	return count, nil
}
