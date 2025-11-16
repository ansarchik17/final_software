package bridge

type IPricingStrategy interface {
	BasePricePerDay() int
}
