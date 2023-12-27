package vehicle

import "errors"

type AccessoryName string

type AccessoryDetails struct {
	Description string
	Price       float32
}

type VehicleAccessories map[AccessoryName]AccessoryDetails

var (
	ErrInvalidAccessoryPrice  = errors.New("accessory price must be positive")
	ErrNoAccessoryDescription = errors.New("empty description")
)
