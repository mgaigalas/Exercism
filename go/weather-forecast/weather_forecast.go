// Package weather provides tools for
// weather forecast.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast returns forecast string value containing city and forecast condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
