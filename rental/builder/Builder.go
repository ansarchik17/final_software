package builder

import (
	"final_software/rental/bridge"
	"final_software/rental/observer/subject"
	"final_software/rental/pricing/decorator"
	"final_software/rental/vehicle"
)

type RentalBuilder struct {
	contract *RentalContract
	events   subject.ISubject
}

var _ IRentalBuilder = (*RentalBuilder)(nil)

func NewRentalBuilder(events subject.ISubject) *RentalBuilder {
	return &RentalBuilder{
		contract: &RentalContract{},
		events:   events,
	}
}

func (b *RentalBuilder) SetClient(name string) IRentalBuilder {
	b.contract.ClientName = name
	return b
}

func (b *RentalBuilder) SetBranch(name string) IRentalBuilder {
	b.contract.Branch = name
	return b
}

func (b *RentalBuilder) SetVehicle(car vehicle.IVehicle) IRentalBuilder {
	b.contract.Vehicle = car

	// Create Bridge strategy
	strategy := bridge.NewVehiclePricing(car)

	// Create initial pricer (Bridge → Decorator adapter)
	b.contract.Pricer = bridge.NewBridgePricer(car.TypeName(), strategy)

	return b
}

func (b *RentalBuilder) SetDays(days int) IRentalBuilder {
	b.contract.Days = days
	return b
}

func (b *RentalBuilder) SetPayment(method string) IRentalBuilder {
	b.contract.PaymentType = method
	return b
}

func (b *RentalBuilder) SetPricer(p decorator.IPricer) IRentalBuilder {
	b.contract.Pricer = p
	return b
}

func (b *RentalBuilder) Build() *RentalContract {
	b.contract.TotalPrice = b.contract.Pricer.Price(b.contract.Days)

	b.events.Notify("RentalContract created for " + b.contract.ClientName)

	return b.contract
}
