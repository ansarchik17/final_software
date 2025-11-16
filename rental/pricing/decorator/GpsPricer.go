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

func (gpsPrice *GpsPricer) Price(days int) int {
	gpsTotalPrice := gpsPrice.inner.Price(days) + days*gpsPrice.feePerDay
	return gpsTotalPrice
}

func (gpsPrice *GpsPricer) Explain() string {
	return gpsPrice.inner.Explain() + fmt.Sprintf(" + GPS %d ₸/day", gpsPrice.feePerDay)
}
