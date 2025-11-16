package pricing

type IPricer interface {
	Price(days int) int
	Explain() string
}
