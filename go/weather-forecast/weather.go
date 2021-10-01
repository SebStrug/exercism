// Package weather manages weather forecasting.
package weather

var (
	// CurrentCondition contains the type of weather.
	CurrentCondition string
	// CurrentLocation is the position.
	CurrentLocation string
)

// Log returns a string containing the location and weather condition.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
