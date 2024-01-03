package main

import (
	f "factory/vehicle"
	"fmt"
)

func main() {
	vehicle := f.New(
		f.Type(f.SUV),
		f.Model("XLE"),
		f.Color("gray"),
		f.Accessory("Ball Mount", 20., "Cold-forged steel construction."),
		f.Accessory("Cargo Cover", 179., "Retractable cargo cover."),
	)

	fmt.Printf("%+v\n", vehicle)
}
