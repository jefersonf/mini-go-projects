package customvehicles

import "errors"

type VehicleType int

// Vehicle types
const (
	Car VehicleType = iota
	Minivan
	Truck
	Crossover
	SUV
	Electrified
)

var ErrVehicleTypeNotFound = errors.New("vehicle type not found")

func (t VehicleType) String() string {
	switch t {
	case Car, Minivan:
		return "Car or Minivan"
	case Truck:
		return "Truck"
	case Crossover, SUV:
		return "Crossover or SUV"
	case Electrified:
		return "Electrified"
	}
	return "Unknown"
}
