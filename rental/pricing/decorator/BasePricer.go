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

func (p *BasePricer) Price(days int) int {
	return days * p.DailyBase
}

func (p *BasePricer) Explain() string {
	return fmt.Sprintf("%s: %d ₸/day", p.VehicleName, p.DailyBase)
}
