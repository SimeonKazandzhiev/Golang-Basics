package main

import "fmt"

func main(){

	var degrees int
	fmt.Scanln(&degrees)
	var time string
	fmt.Scanln(&time)

	var outfit string
	var shoes string
	
	
	if degrees >= 10 && degrees <=18  {
		switch time {
		case "Morning":
			outfit = "Sweatshirt"
			shoes = "Sneakers"
			break
		case "Afternoon":
			outfit = "Shirt"
			shoes = "Moccasins"
			break
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
			break
		}
	}else if degrees > 18 && degrees <=24{
		switch time {
		case "Morning":
			outfit = "Shirt"
			shoes = "Moccasins"
			break
		case "Afternoon":
			outfit = "T-Shirt"
			shoes = "Sandals"
			break
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
			break
		}
	}else if degrees >= 25{
		switch time {
		case "Morning":
			outfit = "T-Shirt"
			shoes = "Sandals"
			break
		case "Afternoon":
			outfit = "Swim Suit"
			shoes = "Barefoot"
			break
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
			break
		}
	}
	fmt.Printf("It's %d degrees, get your %s and %s.",degrees,outfit,shoes)
}