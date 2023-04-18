package rules_test

import (
	"fmt"

	"github.com/IAmRadek/rules"
)

var isPassengerEconomy = rules.NewAttribute("passengerIsEconomy")
var isPassengerGoldCardHolder = rules.NewAttribute("passengerIsGoldCardHolder")
var isPassengerSilverCardHolder = rules.NewAttribute("passengerIsSilverCardHolder")
var isPassengerDressSmart = rules.NewAttribute("passengerDressIsSmart")
var isPassengerDressCasual = rules.NewAttribute("passengerDressIsCasual")
var baggageWeight = rules.NewVariable[float64]("passengerCarryOnBaggageWeightKg")
var baggageAllowance = rules.NewVariable[float64]("carryOnBaggageAllowanceKg")
var suitableForUpgrade = rules.MustParse(
	"suitableForUpgrade",
	`passengerIsEconomy 
			AND (passengerIsGoldCardHolder OR passengerIsSilverCardHolder) 
			AND (passengerCarryOnBaggageWeightKg LTE carryOnBaggageAllowanceKg) 
			AND passengerDressIsSmart`,
)
var canHaveAdditionalBaggage = rules.MustParse(
	"canHaveAdditionalBaggage",
	`passengerIsEconomy
			AND passengerIsGoldCardHolder
			AND (passengerCarryOnBaggageWeightKg LTE carryOnBaggageAllowanceKg)
			AND passengerDressIsSmart`,
)

func ExampleParse() {
	passengerContext := rules.NewContext(
		isPassengerEconomy(true),
		isPassengerGoldCardHolder(true),
		isPassengerSilverCardHolder(false),
		isPassengerDressSmart(true),
		baggageWeight(4.6),
		baggageAllowance(7),
	)

	result, err := suitableForUpgrade.Evaluate(passengerContext)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// Output: true
}

func ExampleNewContext() {
	globalContext := rules.NewContext(
		baggageAllowance(3),
		isPassengerDressCasual(true),
	)

	passengerContext := rules.NewContext(
		isPassengerEconomy(true),
		isPassengerGoldCardHolder(true),
		isPassengerSilverCardHolder(false),
		isPassengerDressSmart(true),
		baggageWeight(4.6),
		baggageAllowance(7),
	).MergeWith(globalContext)

	fmt.Println(passengerContext)

	result, err := suitableForUpgrade.Evaluate(passengerContext)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// Output:
	// passengerIsEconomy(true), passengerIsGoldCardHolder(true), passengerIsSilverCardHolder(false), passengerDressIsSmart(true), passengerCarryOnBaggageWeightKg(4.6), carryOnBaggageAllowanceKg(7), passengerDressIsCasual(true)
	// true
}

func ExampleRuleSet() {
	ruleSet := rules.NewRuleSet(suitableForUpgrade, canHaveAdditionalBaggage)
	ruleSet.AddOverride(canHaveAdditionalBaggage)

	// Create a context for a passenger
	passengerContext := rules.NewContext(
		isPassengerEconomy(true),
		isPassengerGoldCardHolder(false),
		isPassengerSilverCardHolder(true),
		isPassengerDressSmart(true),
		baggageWeight(4.6),
		baggageAllowance(7),
	)

	// Evaluate the rule set
	result, err := ruleSet.Evaluate(passengerContext)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// Output: true
}
