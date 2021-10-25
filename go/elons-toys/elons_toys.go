package elon

import "fmt"

// Drive updates the number of meters driven based on speed, and
// reduces the battery accordingly
func (car *Car) Drive() {
	car.distance += car.speed
	car.battery -= car.batteryDrain
}

// CanFinish checks if a car is able to drive the race's distance
func (car Car) CanFinish(trackDistance int) bool {
	numMoves := trackDistance / car.speed
	return car.battery >= (numMoves * car.batteryDrain)
}

// DisplayDistance returns the distance displayed on the LED as a string
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery displayed on the LED as a string
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}
