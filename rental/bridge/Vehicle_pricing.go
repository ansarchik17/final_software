package bridge

import "final_software/rental/vehicle"

type VehiclePricingStrategy struct {
	car vehicle.IVehicle
}

var _ IPricingStrategy = (*VehiclePricingStrategy)(nil)

func NewVehiclePricing(car vehicle.IVehicle) *VehiclePricingStrategy {
	return &VehiclePricingStrategy{car: car}
}

func (vehicle *VehiclePricingStrategy) BasePricePerDay() int {
	return vehicle.car.DailyBasePrice()
}
