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

func (visit *MaintenanceVisitor) VisitEconomy(c *vehicle.EconomyCar) {
	visit.Report = append(visit.Report,
		fmt.Sprintf("EconomyCar: check engine, tires, interior (seats=%d)", c.Seats))
}

func (visit *MaintenanceVisitor) VisitSuv(c *vehicle.SuvCar) {
	task := "standard check"
	if c.Awd {
		task = "AWD drivetrain check + standard check"
	}

	visit.Report = append(visit.Report,
		fmt.Sprintf("SUV: %s", task))
}
