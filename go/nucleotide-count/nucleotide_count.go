// Package dna contains utilities for calculation of DNA histograms
package dna

import (
	"errors"
	"unicode"
)

type Histogram map[rune]int

type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	res := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, val := range d {
		val = unicode.ToUpper(val)
		switch val {
		case 'A', 'C', 'G', 'T':
			res[val]++
		default:
			return nil, errors.New("invalid DNA sequence")
		}
	}
	return res, nil
}
