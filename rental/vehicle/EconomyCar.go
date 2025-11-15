package vehicle

type EconomyCar struct {
	DailyBase int
	Seats     int
}

var (
	_ IVehicle   = (*EconomyCar)(nil)
	_ IVisitable = (*EconomyCar)(nil)
)

func (c *EconomyCar) TypeName() string {
	return "EconomyCar"
}

func (c *EconomyCar) DailyBasePrice() int {
	return c.DailyBase
}

func (c *EconomyCar) Accept(v IVisitor) {
	v.VisitEconomy(c)
}
