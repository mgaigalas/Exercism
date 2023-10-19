// Package space is a small library, which calculates age based on planets of origin
package space

const EarthYearInSeconds float64 = 31557600

type Planet string

var earthYearMappings = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
	"Earth":   1.0,
}

// Age returns age in years in Earth based on planet and seconds provided as an arguments
func Age(seconds float64, planet Planet) float64 {
	if earthYears, ok := earthYearMappings[planet]; ok {
		return seconds / EarthYearInSeconds / earthYears
	}
	return -1
}
