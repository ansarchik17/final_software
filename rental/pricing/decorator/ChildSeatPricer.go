package decorator

import (
	"fmt"

	"final_software/rental/pricing"
)

type ChildSeatPricer struct {
	inner     pricing.IPricer
	feePerDay int
}

var _ pricing.IPricer = (*ChildSeatPricer)(nil)

func NewChildSeatPricer(inner pricing.IPricer, fee int) *ChildSeatPricer {
	return &ChildSeatPricer{inner: inner, feePerDay: fee}
}

func (c *ChildSeatPricer) Price(days int) int {
	return c.inner.Price(days) + days*c.feePerDay
}

func (c *ChildSeatPricer) Explain() string {
	return c.inner.Explain() + fmt.Sprintf(" + child seat %d ₸/day", c.feePerDay)
}
