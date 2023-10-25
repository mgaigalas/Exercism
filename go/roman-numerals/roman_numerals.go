// Package romannumerals is a library for converting numeric numbers to Roman numerals
package romannumerals

import "errors"

var singles = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hundreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var thousands = []string{"", "M", "MM", "MMM"}

// ToRomanNumeral converts numeric numbers to Roman numerals
func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("number out of bounds")
	}
	return thousands[input/1000] + hundreds[(input/100)%10] + tens[(input/10)%10] + singles[input%10], nil
}
