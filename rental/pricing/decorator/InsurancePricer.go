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

func (p *InsurancePricer) Price(days int) int {
	return p.inner.Price(days) + days*p.feePerDay
}

func (p *InsurancePricer) Explain() string {
	return p.inner.Explain() + fmt.Sprintf(" + insurance %d ₸/day", p.feePerDay)
}
