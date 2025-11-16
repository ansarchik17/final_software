package factory

import "final_software/rental/vehicle"

type CityBranchFactory struct{}

func NewCityBranchFactory() *CityBranchFactory {
	return &CityBranchFactory{}
}

var _ IVehicleFactory = (*CityBranchFactory)(nil)

func (factory *CityBranchFactory) CreateEconomy() vehicle.IVehicle {
	return &vehicle.EconomyCar{
		DailyBase: 10000,
		Seats:     4,
	}
}

func (factory *CityBranchFactory) CreateSUV() vehicle.IVehicle {
	return &vehicle.SuvCar{
		DailyBase: 20000,
		Awd:       false,
	}
}
