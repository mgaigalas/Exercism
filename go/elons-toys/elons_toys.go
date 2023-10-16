package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery-car.batteryDrain < 0 {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	return (1+(trackDistance-1)/car.speed)*car.batteryDrain <= car.battery
}
