package decorator

import (
	"fmt"

	"final_software/rental/pricing"
)

type InsurancePricer struct {
	inner     pricing.IPricer
	feePerDay int
}

var _ pricing.IPricer = (*InsurancePricer)(nil)

func NewInsurancePricer(inner pricing.IPricer, fee int) *InsurancePricer {
	return &InsurancePricer{inner: inner, feePerDay: fee}
}

func (price *InsurancePricer) Price(days int) int {
	insurancePrice := price.inner.Price(days) + days*price.feePerDay
	return insurancePrice
}

func (price *InsurancePricer) Explain() string {
	return price.inner.Explain() + fmt.Sprintf(" + insurance %d ₸/day", price.feePerDay)
}
