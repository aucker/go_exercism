// Package this is a package for weather
package weather

// CurrentCondition a variable string for currentcondition
var CurrentCondition string
// CurrentLocation variable string to store the currentlocation
var CurrentLocation string

// Forecast function is used to return the current location's weather
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}