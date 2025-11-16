package main

import (
	"fmt"

	"final_software/rental/builder"
	"final_software/rental/factory"
	"final_software/rental/observer/observer"
	"final_software/rental/observer/subject"
	"final_software/rental/pricing/decorator"
	"final_software/rental/vehicle"
	"final_software/rental/visitorimpl"
)

func main() {

	// ======================================================
	// 1. OBSERVER SETUP
	// ======================================================
	fmt.Println("=== OBSERVER SETUP ===")

	events := subject.NewSubject()
	events.Register(&observer.AuditLogger{})
	events.Register(&observer.EmailNotifier{Email: "client@example.com"})

	fmt.Println("Observers registered.\n")

	// ======================================================
	// 2. FACTORY – CREATE VEHICLE
	// ======================================================
	fmt.Println("=== FACTORY DEMO ===")

	var cityFactory factory.IVehicleFactory = factory.NewCityBranchFactory()
	car := cityFactory.CreateEconomy()

	fmt.Println("Factory produced:", car.TypeName(), "price/day:", car.DailyBasePrice())

	// ======================================================
	// 3. BUILDER + BRIDGE + DECORATORS
	// ======================================================
	fmt.Println("\n=== BUILDER + BRIDGE + DECORATOR DEMO ===")

	rb := builder.NewRentalBuilder(events)

	// Step-by-step contract creation using Builder
	rb.SetClient("Ansar").
		SetBranch("City Branch").
		SetVehicle(car). // Bridge pricer is attached automatically
		SetDays(3).
		SetPayment("Card")

	// Now wrap with Decorators (your friend's part)
	basePricer := rb.Build().Pricer // builder initially set BridgePricer

	// Rebuild builder again with updated pricer
	rb = builder.NewRentalBuilder(events)
	contract := rb.
		SetClient("Ansar").
		SetBranch("City Branch").
		SetVehicle(car).
		SetDays(3).
		SetPayment("Card").
		SetPricer(
			decorator.NewGpsPricer(
				decorator.NewChildSeatPricer(
					decorator.NewInsurancePricer(
						basePricer, 2000), // insurance
					1000), // child seat
				1500), // GPS
		).
		Build()

	// Output final contract
	fmt.Println("Contract created:")
	fmt.Println(" Client:", contract.ClientName)
	fmt.Println(" Car:", contract.Vehicle.TypeName())
	fmt.Println(" Days:", contract.Days)
	fmt.Println(" Payment:", contract.PaymentType)
	fmt.Println(" Price explanation:", contract.Pricer.Explain())
	fmt.Println(" Total:", contract.TotalPrice, "₸")

	// ======================================================
	// 4. VISITOR – ADMIN ANALYSIS
	// ======================================================
	fmt.Println("\n=== VISITOR DEMO ===")

	fleet := []vehicle.IVisitable{
		&vehicle.EconomyCar{DailyBase: 10000, Seats: 4},
		&vehicle.SuvCar{DailyBase: 20000, Awd: true},
	}

	// Pricing Visitor
	priceVisitor := visitorimpl.NewPricingVisitor(3)
	for _, car := range fleet {
		car.Accept(priceVisitor)
	}

	fmt.Println("Pricing Analysis:")
	for _, note := range priceVisitor.Notes {
		fmt.Println(" -", note)
	}
	fmt.Println("Total potential:", priceVisitor.Total, "₸")

	// Maintenance Visitor
	maintVisitor := visitorimpl.NewMaintenanceVisitor()
	for _, car := range fleet {
		car.Accept(maintVisitor)
	}

	fmt.Println("\nMaintenance Report:")
	for _, line := range maintVisitor.Report {
		fmt.Println(" -", line)
	}
}
