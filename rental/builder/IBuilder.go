package builder

import (
	"final_software/rental/pricing"
	"final_software/rental/vehicle"
)

type IRentalBuilder interface {
	SetClient(name string) IRentalBuilder
	SetBranch(name string) IRentalBuilder
	SetVehicle(car vehicle.IVehicle) IRentalBuilder
	SetDays(days int) IRentalBuilder
	SetPayment(method string) IRentalBuilder
	SetPricer(price pricing.IPricer) IRentalBuilder
	Build() *RentalContract
}
