package builder

import (
	"final_software/rental/pricing"
	"final_software/rental/vehicle"
)

type RentalContract struct {
	ClientName  string
	Vehicle     vehicle.IVehicle
	Branch      string
	Days        int
	Pricer      pricing.IPricer
	TotalPrice  int
	PaymentType string
}
