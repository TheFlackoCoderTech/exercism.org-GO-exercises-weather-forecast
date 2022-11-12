// Package weather provides tools to convert.
// temperatures to and from weather.
package weather

// CurrentCondition should be written right before the variable that it is describing, start with its name, and end with a period.
var CurrentCondition string
// CurrentLocation should be written right before the variable that it is describing, start with its name, and end with a period.
var CurrentLocation string

// Forecast returns an value to city & condition string's.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}