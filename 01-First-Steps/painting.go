package main

import "fmt"


func main() {

	var protection int
	fmt.Scanln(&protection)
	var paintNeeded int
	fmt.Scanln(&paintNeeded)
	var thinner int
	fmt.Scanln(&thinner)
	var hoursNeeded int
	fmt.Scanln(&hoursNeeded)


	var protectionPrice float32 = (float32(protection) + 2) * 1.5
	var paintPrice float32 = float32(paintNeeded) * 1.1 * 14.50
	var thinnerPrice float32 = float32(thinner) * 5.00
	var totalPriceMaterials float32 = protectionPrice + paintPrice + thinnerPrice + 0.40
	var layborsPrice float32 = totalPriceMaterials * 0.3 * 8
	var totalPrice float32 = totalPriceMaterials + layborsPrice

	fmt.Println(totalPrice)
}