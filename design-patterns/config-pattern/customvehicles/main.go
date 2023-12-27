package main

import (
	cv "customvehicles/factory"
	"fmt"
)

func main() {
	vehicle := cv.NewVehicle(
		cv.Type(cv.SUV),
		cv.Model("XLE"),
		cv.Color("gray"),
		cv.Accessory("Ball Mount", 20., "Cold-forged steel construction."),
		cv.Accessory("Cargo Cover", 179., "Retractable cargo cover."),
	)

	fmt.Printf("%+v\n", vehicle)
}
