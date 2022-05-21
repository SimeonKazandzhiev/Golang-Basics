package main

import "fmt"

func main(){

	var finalPrice float32
	
	var product string
	fmt.Scanln(&product)
	var town string
	fmt.Scanln(&town)
	var quantity float32
	fmt.Scanln(&quantity)

	if town == "Sofia"{
		switch product{
		case "coffee": finalPrice = quantity * 0.50
		case "water": finalPrice = quantity * 0.80
		case "beer": finalPrice = quantity * 1.20
		case "sweets": finalPrice = quantity * 1.45
		case "peanuts": finalPrice = quantity * 1.60
		}
	}else if town == "Plovdiv"{
		switch product{
			case "coffee": finalPrice = quantity * 0.40
			case "water": finalPrice = quantity * 0.70
			case "beer": finalPrice = quantity * 1.15
			case "sweets": finalPrice = quantity * 1.30
			case "peanuts": finalPrice = quantity * 1.50
			}
	}else if town == "Varna"{
		switch product{
			case "coffee": finalPrice = quantity * 0.45
			case "water": finalPrice = quantity * 0.70
			case "beer": finalPrice = quantity * 1.10
			case "sweets": finalPrice = quantity * 1.35
			case "peanuts": finalPrice = quantity * 1.55
			}
	}

	fmt.Println(finalPrice)

}