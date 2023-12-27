package vehicle

import "errors"

type VehicleModel string

// Vehicle models
const (
	// Car & minivan
	XLE      VehicleModel = "XLE"
	Limited  VehicleModel = "Limited"
	Platinum VehicleModel = "Platinum"
	// truck
	SR           VehicleModel = "SR"
	TRDPreRunner VehicleModel = "TRD PreRunner"
	TRDOffRoad   VehicleModel = "TRD Off-Road"
	// crossover & SUV
	LE             VehicleModel = "LE"
	HybridLE       VehicleModel = "Hybrid LE"
	XLW            VehicleModel = "XLE"
	HybridPlatinum VehicleModel = "Hybrid Platinum"
)

var ErrModelNotAvailable = errors.New("model not available")
