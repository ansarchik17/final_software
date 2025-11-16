package factory

import "final_software/rental/vehicle"

type AirportBranchFactory struct{}

func NewAirportBranchFactory() *AirportBranchFactory {
	return &AirportBranchFactory{}
}

var _ IVehicleFactory = (*AirportBranchFactory)(nil)

func (factory *AirportBranchFactory) CreateEconomy() vehicle.IVehicle {
	return &vehicle.EconomyCar{
		DailyBase: 12000,
		Seats:     4,
	}
}

func (factory *AirportBranchFactory) CreateSUV() vehicle.IVehicle {
	return &vehicle.SuvCar{
		DailyBase: 25000,
		Awd:       true,
	}
}
