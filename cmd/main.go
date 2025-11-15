package main

import (
	"fmt"

	"final_software/rental/factory"
	"final_software/rental/pricing/decorator"
	"final_software/rental/vehicle"
	"final_software/rental/visitorimpl"
)

func main() {
	fmt.Println("=== FACTORY DEMO ===")

	cityFactory := factory.NewCityBranchFactory()
	airportFactory := factory.NewAirportBranchFactory()

	cityEco := cityFactory.CreateEconomy()
	airEco := airportFactory.CreateEconomy()

	fmt.Println("City Economy:", cityEco.TypeName(), "price/day:", cityEco.DailyBasePrice())
	fmt.Println("Airport Economy:", airEco.TypeName(), "price/day:", airEco.DailyBasePrice())

	// ------------------------------
	fmt.Println("\n=== DECORATOR DEMO ===")

	base := decorator.NewBasePricer("EconomyCar", 10000)

	withInsurance := decorator.NewInsurancePricer(base, 2000)
	withSeat := decorator.NewChildSeatPricer(withInsurance, 1000)
	withGps := decorator.NewGpsPricer(withSeat, 1500)

	fmt.Println("Explain:", withGps.Explain())
	fmt.Println("Price for 3 days:", withGps.Price(3), "₸")

	// ------------------------------
	fmt.Println("\n=== VISITOR DEMO ===")

	// fleet for visitor = Visitable elements
	fleet := []vehicle.IVisitable{
		&vehicle.EconomyCar{DailyBase: 10000, Seats: 4},
		&vehicle.SuvCar{DailyBase: 20000, Awd: true},
	}

	priceVisitor := visitorimpl.NewPricingVisitor(3)
	for _, car := range fleet {
		car.Accept(priceVisitor)
	}

	fmt.Println("Visitor Pricing Notes:")
	for _, note := range priceVisitor.Notes {
		fmt.Println(" -", note)
	}
	fmt.Println("Total potential:", priceVisitor.Total, "₸")

	maintVisitor := visitorimpl.NewMaintenanceVisitor()
	for _, car := range fleet {
		car.Accept(maintVisitor)
	}

	fmt.Println("\nVisitor Maintenance:")
	for _, line := range maintVisitor.Report {
		fmt.Println(" -", line)
	}
}
