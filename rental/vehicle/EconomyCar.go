package vehicle

type EconomyCar struct {
	DailyBase int
	Seats     int
}

var (
	_ IVehicle   = (*EconomyCar)(nil)
	_ IVisitable = (*EconomyCar)(nil)
)

func (car *EconomyCar) TypeName() string {
	return "EconomyCar"
}

func (car *EconomyCar) DailyBasePrice() int {
	return car.DailyBase
}

func (car *EconomyCar) Accept(v IVisitor) {
	v.VisitEconomy(car)
}
