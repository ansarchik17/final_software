package decorator

import "fmt"

type GpsPricer struct {
	inner     IPricer
	feePerDay int
}

var _ IPricer = (*GpsPricer)(nil)

func NewGpsPricer(inner IPricer, fee int) *GpsPricer {
	return &GpsPricer{inner: inner, feePerDay: fee}
}

func (p *GpsPricer) Price(days int) int {
	return p.inner.Price(days) + days*p.feePerDay
}

func (p *GpsPricer) Explain() string {
	return p.inner.Explain() + fmt.Sprintf(" + GPS %d ₸/day", p.feePerDay)
}
