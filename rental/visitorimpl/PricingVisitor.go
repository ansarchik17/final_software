package visitorimpl

import (
	"fmt"

	"final_software/rental/vehicle"
)

type PricingVisitor struct {
	Days  int
	Total int
	Notes []string
}

var _ vehicle.IVisitor = (*PricingVisitor)(nil)

func NewPricingVisitor(days int) *PricingVisitor {
	return &PricingVisitor{Days: days}
}

func (visit *PricingVisitor) VisitEconomy(c *vehicle.EconomyCar) {
	price := visit.Days * c.DailyBase
	visit.Total += price
	visit.Notes = append(visit.Notes,
		fmt.Sprintf("Economy × %d days = %d ₸", visit.Days, price))
}

func (visit *PricingVisitor) VisitSuv(c *vehicle.SuvCar) {
	base := visit.Days * c.DailyBase
	if c.Awd {
		base = int(float64(base) * 1.10)
	}

	visit.Total += base
	visit.Notes = append(visit.Notes,
		fmt.Sprintf("SUV × %d days (AWD=%v) = %d ₸", visit.Days, c.Awd, base))
}
