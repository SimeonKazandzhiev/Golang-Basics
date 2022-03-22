package main

import "fmt"

func main() {

	var budget float32
	fmt.Scanln(&budget)
	var season string
	fmt.Scanln(&season)
	var fishersCount int
	fmt.Scanln(&fishersCount)

	var finalPrice float32

	if season == "Spring" {
		finalPrice = 3000
	} else if season == "Summer" || season == "Autumn" {
		finalPrice = 4200
	} else if season == "Winter" {
		finalPrice = 2600
	}

	if fishersCount <= 6 {
		finalPrice = finalPrice * 0.9
	} else if fishersCount > 6 && fishersCount <= 11 {
		finalPrice = finalPrice * 0.85
	} else {
		finalPrice = finalPrice * 0.75
	}

	if fishersCount%2 == 0 && season != "Autumn" {
		finalPrice = finalPrice * 0.95
	}

	if budget >= finalPrice {
		fmt.Printf("Yes! You have %.2f leva left.", budget-finalPrice)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva.", finalPrice-budget)
	}
}
