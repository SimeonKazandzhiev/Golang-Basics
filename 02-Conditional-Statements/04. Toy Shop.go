package main

import "fmt"

func main () {
	
	var tripPrice float32
	fmt.Scanln(&tripPrice) 

	var puzzlesCount, dollsCount, bearsCount, minionsCount, trucksCount int
	fmt.Scanln(&puzzlesCount) 
	fmt.Scanln(&dollsCount) 
	fmt.Scanln(&bearsCount) 
	fmt.Scanln(&minionsCount)
	fmt.Scanln(&trucksCount) 

	var profit float32 = (float32(puzzlesCount) * 2.60) + (float32(dollsCount) * 3.00) + (float32(bearsCount) * 4.10) + (float32(minionsCount) * 8.20) + (float32(trucksCount) * 2.00) 
	var countToys int = puzzlesCount + dollsCount + bearsCount + minionsCount + trucksCount 

	if countToys >= 50 {
		profit = profit - 0.25 * profit 
	}

	var rent float32 = 0.1 * profit 
	var finalSum float32 = profit - rent

	if tripPrice <= finalSum {
		var leftMoney float32 = finalSum - tripPrice 
		fmt.Printf("Yes! %.2f lv left.", leftMoney)
	} else {
		var needMoney float32 = tripPrice - finalSum
		fmt.Printf("Not enough money! %.2f lv needed.", needMoney)
	}

}