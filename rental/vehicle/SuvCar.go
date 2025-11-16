package vehicle

type SuvCar struct {
	DailyBase int
	Awd       bool
}

var (
	_ IVehicle   = (*SuvCar)(nil)
	_ IVisitable = (*SuvCar)(nil)
)

func (car *SuvCar) TypeName() string {
	return "SuvCar"
}

func (car *SuvCar) DailyBasePrice() int {
	return car.DailyBase
}

func (car *SuvCar) Accept(v IVisitor) {
	v.VisitSuv(car)
}
