package decorator

import "fmt"

type InsurancePricer struct {
	inner     IPricer
	feePerDay int
}

var _ IPricer = (*InsurancePricer)(nil)

func NewInsurancePricer(inner IPricer, fee int) *InsurancePricer {
	return &InsurancePricer{inner: inner, feePerDay: fee}
}

func (price *InsurancePricer) Price(days int) int {
	insurancePrice := price.inner.Price(days) + days*price.feePerDay
	return insurancePrice
}

func (price *InsurancePricer) Explain() string {
	return price.inner.Explain() + fmt.Sprintf(" + insurance %d ₸/day", price.feePerDay)
}
