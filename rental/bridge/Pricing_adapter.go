package bridge

import (
	"fmt"

	"final_software/rental/pricing"
)

type BridgePricer struct {
	vehicleName string
	strategy    IPricingStrategy
}

var _ pricing.IPricer = (*BridgePricer)(nil)

func NewBridgePricer(name string, s IPricingStrategy) *BridgePricer {
	return &BridgePricer{
		vehicleName: name,
		strategy:    s,
	}
}

func (price *BridgePricer) Price(days int) int {
	return price.strategy.BasePricePerDay() * days
}

func (price *BridgePricer) Explain() string {
	return fmt.Sprintf("%s: %d ₸/day", price.vehicleName, price.strategy.BasePricePerDay())
}
