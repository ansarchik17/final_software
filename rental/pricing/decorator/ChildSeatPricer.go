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

func (childSeatPrice *ChildSeatPricer) Price(days int) int {
	totalChildSeatPrice := childSeatPrice.inner.Price(days) + days*childSeatPrice.feePerDay
	return totalChildSeatPrice
}

func (childSeatPrice *ChildSeatPricer) Explain() string {
	return childSeatPrice.inner.Explain() + fmt.Sprintf(" + child seat %d ₸/day", childSeatPrice.feePerDay)
}
