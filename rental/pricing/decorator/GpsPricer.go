package decorator

import (
	"fmt"

	"final_software/rental/pricing"
)

type GpsPricer struct {
	inner     pricing.IPricer
	feePerDay int
}

var _ pricing.IPricer = (*GpsPricer)(nil)

func NewGpsPricer(inner pricing.IPricer, fee int) *GpsPricer {
	return &GpsPricer{inner: inner, feePerDay: fee}
}

func (gpsPrice *GpsPricer) Price(days int) int {
	gpsTotalPrice := gpsPrice.inner.Price(days) + days*gpsPrice.feePerDay
	return gpsTotalPrice
}

func (gpsPrice *GpsPricer) Explain() string {
	return gpsPrice.inner.Explain() + fmt.Sprintf(" + GPS %d ₸/day", gpsPrice.feePerDay)
}
