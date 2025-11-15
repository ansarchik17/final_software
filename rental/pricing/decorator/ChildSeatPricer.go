package decorator

import "fmt"

type ChildSeatPricer struct {
	inner     IPricer
	feePerDay int
}

var _ IPricer = (*ChildSeatPricer)(nil)

func NewChildSeatPricer(inner IPricer, fee int) *ChildSeatPricer {
	return &ChildSeatPricer{inner: inner, feePerDay: fee}
}

func (p *ChildSeatPricer) Price(days int) int {
	return p.inner.Price(days) + days*p.feePerDay
}

func (p *ChildSeatPricer) Explain() string {
	return p.inner.Explain() + fmt.Sprintf(" + child seat %d ₸/day", p.feePerDay)
}
