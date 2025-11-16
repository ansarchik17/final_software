package bridge

import (
	"final_software/rental/pricing/decorator"
	"fmt"
)

type BridgePricer struct {
	vehicleName string
	strategy    IPricingStrategy
}

var _ decorator.IPricer = (*BridgePricer)(nil)

func NewBridgePricer(name string, s IPricingStrategy) *BridgePricer {
	return &BridgePricer{
		vehicleName: name,
		strategy:    s,
	}
}

func (p *BridgePricer) Price(days int) int {
	return p.strategy.BasePricePerDay() * days
}

func (p *BridgePricer) Explain() string {
	return fmt.Sprintf("%s: %d ₸/day", p.vehicleName, p.strategy.BasePricePerDay())
}
