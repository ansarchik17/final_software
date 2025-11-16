package main

import (
	"fmt"

	"final_software/rental/bridge"
	"final_software/rental/builder"
	"final_software/rental/factory"
	"final_software/rental/observer/observer"
	"final_software/rental/observer/subject"
	"final_software/rental/pricing"
	"final_software/rental/pricing/decorator"
	"final_software/rental/vehicle"
	"final_software/rental/visitorimpl"
)

func main() {
	// Abstract Factory
	cityBranchVehicleFactoryFromAbstractFactoryPattern := factory.NewCityBranchFactory()
	airportBranchVehicleFactoryFromAbstractFactoryPattern := factory.NewAirportBranchFactory()

	cityBranchEconomyVehicleFromAbstractFactoryPattern := cityBranchVehicleFactoryFromAbstractFactoryPattern.CreateEconomy()
	cityBranchSuvVehicleFromAbstractFactoryPattern := cityBranchVehicleFactoryFromAbstractFactoryPattern.CreateSUV()
	airportBranchEconomyVehicleFromAbstractFactoryPattern := airportBranchVehicleFactoryFromAbstractFactoryPattern.CreateEconomy()
	airportBranchSuvVehicleFromAbstractFactoryPattern := airportBranchVehicleFactoryFromAbstractFactoryPattern.CreateSUV()

	fmt.Println("===== Abstract Factory: vehicles created for different rental branches =====")
	printVehicleDetailsFromAbstractFactoryPattern := func(branchNameForVehicleFromAbstractFactoryPattern string, createdVehicleFromAbstractFactoryPattern vehicle.IVehicle) {
		fmt.Printf(
			"Branch=%s, VehicleType=%s, DailyBasePrice=%d\n",
			branchNameForVehicleFromAbstractFactoryPattern,
			createdVehicleFromAbstractFactoryPattern.TypeName(),
			createdVehicleFromAbstractFactoryPattern.DailyBasePrice(),
		)
	}

	printVehicleDetailsFromAbstractFactoryPattern("City branch (Economy)", cityBranchEconomyVehicleFromAbstractFactoryPattern)
	printVehicleDetailsFromAbstractFactoryPattern("City branch (SUV)", cityBranchSuvVehicleFromAbstractFactoryPattern)
	printVehicleDetailsFromAbstractFactoryPattern("Airport branch (Economy)", airportBranchEconomyVehicleFromAbstractFactoryPattern)
	printVehicleDetailsFromAbstractFactoryPattern("Airport branch (SUV)", airportBranchSuvVehicleFromAbstractFactoryPattern)

	// Bridge
	fmt.Println()
	fmt.Println("===== Bridge: base pricing for city economy car =====")

	cityBranchEconomyVehiclePricingStrategyFromBridgePattern := bridge.NewVehiclePricing(cityBranchEconomyVehicleFromAbstractFactoryPattern)
	cityBranchEconomyVehicleBridgePricerFromBridgePattern := bridge.NewBridgePricer(
		cityBranchEconomyVehicleFromAbstractFactoryPattern.TypeName(),
		cityBranchEconomyVehiclePricingStrategyFromBridgePattern,
	)

	totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns := 3

	fmt.Printf(
		"Base pricing from Bridge pattern for %d days: %d ₸ (%s)\n",
		totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns,
		cityBranchEconomyVehicleBridgePricerFromBridgePattern.Price(totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns),
		cityBranchEconomyVehicleBridgePricerFromBridgePattern.Explain(),
	)

	// Decorator
	fmt.Println()
	fmt.Println("===== Decorator: add-ons for city economy car pricing =====")

	var decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern pricing.IPricer = cityBranchEconomyVehicleBridgePricerFromBridgePattern
	decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern = decorator.NewInsurancePricer(decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern, 2000)
	decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern = decorator.NewGpsPricer(decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern, 500)
	decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern = decorator.NewChildSeatPricer(decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern, 1000)

	fmt.Printf(
		"Pricing explanation from Decorator pattern: %s\n",
		decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern.Explain(),
	)
	fmt.Printf(
		"Total price for %d days from Bridge and Decorator patterns combined: %d ₸\n",
		totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns,
		decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern.Price(totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns),
	)

	// Observer
	fmt.Println()
	fmt.Println("===== Observer: observers attached to rental events subject =====")

	rentalEventSubjectForObserverPattern := subject.NewSubject()

	emailNotifierObserverForRentalEventsFromObserverPattern := &observer.EmailNotifier{Email: "customer@example.com"}
	auditLoggerObserverForRentalEventsFromObserverPattern := &observer.AuditLogger{}

	rentalEventSubjectForObserverPattern.Register(emailNotifierObserverForRentalEventsFromObserverPattern)
	rentalEventSubjectForObserverPattern.Register(auditLoggerObserverForRentalEventsFromObserverPattern)

	rentalEventSubjectForObserverPattern.Notify("Initial system startup event before Builder pattern usage")

	// Builder
	fmt.Println()
	fmt.Println("===== Builder: rental contract assembled with integrated Bridge, Decorator and Observer =====")

	rentalContractBuilderWithObserverAndBridgeAndDecoratorPatterns := builder.NewRentalBuilder(rentalEventSubjectForObserverPattern)

	rentalContractBuiltByBuilderPatternWithDecoratedPricer := rentalContractBuilderWithObserverAndBridgeAndDecoratorPatterns.
		SetClient("Sample Customer From Builder Pattern").
		SetBranch("City branch from Builder Pattern").
		SetVehicle(cityBranchEconomyVehicleFromAbstractFactoryPattern).
		SetDays(totalRentalDaysForPricingDemonstrationFromBridgeAndDecoratorPatterns).
		SetPayment("Card payment from Builder Pattern").
		SetPricer(decoratedVehiclePricerForCityBranchEconomyCarFromDecoratorPattern).
		Build()

	fmt.Printf(
		"Rental contract built by Builder pattern (with Bridge and Decorator inside): %+v\n",
		rentalContractBuiltByBuilderPatternWithDecoratedPricer,
	)

	rentalEventSubjectForObserverPattern.Notify("Rental contract fully processed after Builder pattern usage")

	// Visitor
	fmt.Println()
	fmt.Println("===== Visitor: pricing and maintenance visitors over entire vehicle fleet =====")

	totalRentalDaysForVisitorPricingDemonstrationFromVisitorPattern := 5
	pricingVisitorForEntireVehicleFleetFromVisitorPattern := visitorimpl.NewPricingVisitor(totalRentalDaysForVisitorPricingDemonstrationFromVisitorPattern)
	maintenanceVisitorForEntireVehicleFleetFromVisitorPattern := visitorimpl.NewMaintenanceVisitor()

	visitableVehiclesForVisitorPattern := []vehicle.IVisitable{}

	if visitableCityBranchEconomyVehicleFromVisitorPattern, ok := cityBranchEconomyVehicleFromAbstractFactoryPattern.(vehicle.IVisitable); ok {
		visitableVehiclesForVisitorPattern = append(visitableVehiclesForVisitorPattern, visitableCityBranchEconomyVehicleFromVisitorPattern)
	}
	if visitableCityBranchSuvVehicleFromVisitorPattern, ok := cityBranchSuvVehicleFromAbstractFactoryPattern.(vehicle.IVisitable); ok {
		visitableVehiclesForVisitorPattern = append(visitableVehiclesForVisitorPattern, visitableCityBranchSuvVehicleFromVisitorPattern)
	}
	if visitableAirportBranchEconomyVehicleFromVisitorPattern, ok := airportBranchEconomyVehicleFromAbstractFactoryPattern.(vehicle.IVisitable); ok {
		visitableVehiclesForVisitorPattern = append(visitableVehiclesForVisitorPattern, visitableAirportBranchEconomyVehicleFromVisitorPattern)
	}
	if visitableAirportBranchSuvVehicleFromVisitorPattern, ok := airportBranchSuvVehicleFromAbstractFactoryPattern.(vehicle.IVisitable); ok {
		visitableVehiclesForVisitorPattern = append(visitableVehiclesForVisitorPattern, visitableAirportBranchSuvVehicleFromVisitorPattern)
	}

	for _, visitableVehicleFromEntireFleetForVisitorPattern := range visitableVehiclesForVisitorPattern {
		visitableVehicleFromEntireFleetForVisitorPattern.Accept(pricingVisitorForEntireVehicleFleetFromVisitorPattern)
		visitableVehicleFromEntireFleetForVisitorPattern.Accept(maintenanceVisitorForEntireVehicleFleetFromVisitorPattern)
	}

	fmt.Printf(
		"Visitor pricing total for %d days over entire vehicle fleet: %d ₸\n",
		totalRentalDaysForVisitorPricingDemonstrationFromVisitorPattern,
		pricingVisitorForEntireVehicleFleetFromVisitorPattern.Total,
	)

	fmt.Println("Visitor pricing notes for entire vehicle fleet from Visitor pattern:")
	for _, pricingNoteFromVisitorForEntireVehicleFleet := range pricingVisitorForEntireVehicleFleetFromVisitorPattern.Notes {
		fmt.Println(" -", pricingNoteFromVisitorForEntireVehicleFleet)
	}

	fmt.Println("Visitor maintenance report for entire vehicle fleet from Visitor pattern:")
	for _, maintenanceReportLineFromVisitorForEntireVehicleFleet := range maintenanceVisitorForEntireVehicleFleetFromVisitorPattern.Report {
		fmt.Println(" -", maintenanceReportLineFromVisitorForEntireVehicleFleet)
	}
}
