package decorator

import "fmt"

// TODO поменять на бридж ансара
type BasePricer struct {
	VehicleName string
	DailyBase   int
}

var _ IPricer = (*BasePricer)(nil)

func NewBasePricer(name string, base int) *BasePricer {
	return &BasePricer{
		VehicleName: name,
		DailyBase:   base,
	}
}

func (basePrice *BasePricer) Price(days int) int {
	baseprice := days * basePrice.DailyBase
	return baseprice
}

func (basePrice *BasePricer) Explain() string {
	return fmt.Sprintf("%s: %d ₸/day", basePrice.VehicleName, basePrice.DailyBase)
}
