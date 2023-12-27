package vehicle

import "log"

type Vehicle struct {
	VehicleOptions
}

func NewVehicle(opts ...Option) *Vehicle {
	options := defaultOptions()
	for _, customOption := range opts {
		err := customOption(&options)
		if err != nil {
			log.Println(err)
		}
	}
	return &Vehicle{options}
}
