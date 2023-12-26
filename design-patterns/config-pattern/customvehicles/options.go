package customvehicles

import "fmt"

type Option func(*VehicleOptions) error

type VehicleOptions struct {
	Type        VehicleType
	Model       VehicleModel
	Color       VehicleColor
	Accessories VehicleAccessories
}

func Type(vehicleType VehicleType) Option {
	return func(vo *VehicleOptions) (err error) {
		switch vehicleType {
		case Car, Minivan, Truck, Crossover, SUV, Electrified:
			vo.Type = vehicleType
		default:
			err = ErrVehicleTypeNotFound
		}
		return
	}
}

func Model(model VehicleModel) Option {
	return func(vo *VehicleOptions) (err error) {
		switch model {
		case XLE, Limited, Platinum, SR, TRDPreRunner, TRDOffRoad, LE, HybridLE, HybridPlatinum:
			vo.Model = model
		default:
			err = fmt.Errorf("%w: %s", ErrModelNotAvailable, model)
		}
		return
	}
}

func Color(color string) Option {
	return func(vo *VehicleOptions) error {
		c, err := fromName(color)
		if err != nil {
			return fmt.Errorf("%w: %s", err, color)
		}
		vo.Color = c
		return nil
	}
}

func Accessory(name string, price float32, description string) Option {
	return func(vo *VehicleOptions) error {
		// TODO: validate accessory avaialability here
		if price <= 0 {
			return fmt.Errorf("%w: %v", ErrInvalidAccessoryPrice, price)
		}
		if len(description) == 0 {
			return fmt.Errorf("%s: %w", name, ErrNoAccessoryDescription)
		}
		vo.Accessories[AccessoryName(name)] = AccessoryDetails{
			Price:       price,
			Description: description,
		}
		return nil
	}
}

func defaultOptions() VehicleOptions {
	return VehicleOptions{
		Type:        Car,
		Model:       XLE,
		Color:       Black,
		Accessories: make(map[AccessoryName]AccessoryDetails),
	}
}
