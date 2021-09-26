// Package lasagna is the solution to
// https://exercism.org/tracks/go/exercises/weather-forecast
package weather

// CurrentLocation is variable where condition info state.
var CurrentCondition string

// CurrentLocation is variable where city info state.
var CurrentLocation string

// Forecast returns the weather to spesific city and condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
