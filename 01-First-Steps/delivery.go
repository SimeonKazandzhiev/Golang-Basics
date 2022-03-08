package main

import "fmt"

func main() {

	var chickenMenus int
	fmt.Scanln(&chickenMenus)
	var fishMenus int
	fmt.Scanln(&fishMenus)
	var vegMenus int
	fmt.Scanln(&vegMenus)

	var totalPriceForMenus float32 = 
	(float32(chickenMenus) * 10.35) + (float32(fishMenus) * 12.40) + (float32(vegMenus) * 8.15)

	var dessertPrice float32 = totalPriceForMenus * 0.2

	var finalPrice = totalPriceForMenus + dessertPrice + 2.50

	fmt.Println(finalPrice)
}