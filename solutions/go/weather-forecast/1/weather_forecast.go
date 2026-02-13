// Package weather is used to furnish forecast.
package weather

var (
	// CurrentCondition represent Current condition Location.
	CurrentCondition string
	// CurrentLocation represent Current condition.
	CurrentLocation string
)

// Forecast takes in city and condtion as string arguments
// and returns a string contactinating the location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
