package vehicle

type IVisitor interface {
	VisitEconomy(*EconomyCar)
	VisitSuv(*SuvCar)
}
