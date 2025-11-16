package builder

import (
	"final_software/rental/pricing/decorator"
	"final_software/rental/vehicle"
)

type RentalContract struct {
	ClientName  string
	Vehicle     vehicle.IVehicle
	Branch      string
	Days        int
	Pricer      decorator.IPricer
	TotalPrice  int
	PaymentType string
}
