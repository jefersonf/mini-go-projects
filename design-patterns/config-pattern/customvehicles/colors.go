package customvehicles

import (
	"errors"
	"strings"
)

type VehicleColor int

// Vehicle colors
const (
	Black VehicleColor = iota
	MagneticGrayMetallic
	OxygenWhite
	HeavyMetal
	SupersonicRed
)

var ErrColorNotAvailable = errors.New("color not available")

func fromName(colorName string) (color VehicleColor, err error) {
	switch strings.ToLower(colorName) {
	case "black":
		color = Black
	case "magnetic gray metallic", "gray metallic", "gray":
		color = MagneticGrayMetallic
	case "oxygen white", "white":
		color = OxygenWhite
	case "heavy metal", "metal":
		color = HeavyMetal
	case "supersonic red", "red":
		color = SupersonicRed
	default:
		err = errors.New("color not available")
	}
	return color, err
}

func (c VehicleColor) String() (color string) {
	switch c {
	case Black:
		color = "Black"
	case MagneticGrayMetallic:
		color = "Magnetic Gray Metallic"
	case OxygenWhite:
		color = "Oxygen White"
	case HeavyMetal:
		color = "Heavy Metal"
	case SupersonicRed:
		color = "Supersonic Red"
	default:
		color = "Unknown"
	}
	return
}
