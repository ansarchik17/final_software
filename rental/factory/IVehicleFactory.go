package factory

import "final_software/rental/vehicle"

type IVehicleFactory interface {
	CreateEconomy() vehicle.IVehicle
	CreateSUV() vehicle.IVehicle
}
