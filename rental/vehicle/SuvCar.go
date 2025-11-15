package vehicle

type SuvCar struct {
	DailyBase int
	Awd       bool
}

var (
	_ IVehicle   = (*SuvCar)(nil)
	_ IVisitable = (*SuvCar)(nil)
)

func (c *SuvCar) TypeName() string {
	return "SuvCar"
}

func (c *SuvCar) DailyBasePrice() int {
	return c.DailyBase
}

func (c *SuvCar) Accept(v IVisitor) {
	v.VisitSuv(c)
}
