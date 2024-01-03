package vehicle

import "log"

// Vehicle represets an vehicle that could have costum options.
type Vehicle struct {
	VehicleOptions
}

// New creates a new vehicle that can be customized by input opti meons.
func New(opts ...Option) *Vehicle {
	options := defaultOptions()
	for _, customOption := range opts {
		err := customOption(&options)
		if err != nil {
			log.Println(err)
		}
	}
	return &Vehicle{options}
}
