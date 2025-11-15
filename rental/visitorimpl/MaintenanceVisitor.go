package visitorimpl

import (
	"fmt"

	"final_software/rental/vehicle"
)

type MaintenanceVisitor struct {
	Report []string
}

var _ vehicle.IVisitor = (*MaintenanceVisitor)(nil)

func NewMaintenanceVisitor() *MaintenanceVisitor {
	return &MaintenanceVisitor{}
}

func (v *MaintenanceVisitor) VisitEconomy(c *vehicle.EconomyCar) {
	v.Report = append(v.Report,
		fmt.Sprintf("EconomyCar: check engine, tires, interior (seats=%d)", c.Seats))
}

func (v *MaintenanceVisitor) VisitSuv(c *vehicle.SuvCar) {
	task := "standard check"
	if c.Awd {
		task = "AWD drivetrain check + standard check"
	}

	v.Report = append(v.Report,
		fmt.Sprintf("SUV: %s", task))
}
