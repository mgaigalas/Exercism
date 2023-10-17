// Package raindrops contains tools for converting numbers to rain sounds
package raindrops

import "strconv"

type SoundEntry struct {
	Number int
	Sound  string
}

// Convert converts number to rain sounds
func Convert(number int) string {
	// can't use maps as we need ordered collection
	dictionary := []SoundEntry{
		{
			Number: 3,
			Sound:  "Pling",
		},
		{
			Number: 5,
			Sound:  "Plang",
		},
		{
			Number: 7,
			Sound:  "Plong",
		},
	}

	var res string
	for _, val := range dictionary {
		if number%val.Number == 0 {
			res += val.Sound
		}
	}

	if res == "" {
		return strconv.Itoa(number)
	}
	return res
}
