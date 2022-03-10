package main

import "fmt"


func main() {

	var pens int
	fmt.Scanln(&pens)
	var markers int
	fmt.Scanln(&markers)
	var litersOfCleaningLiquid int
	fmt.Scanln(&litersOfCleaningLiquid)
	var discount int
	fmt.Scanln(&discount)

	var totalPensPrice float32 = float32(pens) * 5.80
	var totalMarkersPrice float32 = float32(markers) * 7.20
	var totalLiquidPrice float32 = float32(litersOfCleaningLiquid) * 1.20

	var totalPrice float32 = totalPensPrice + totalLiquidPrice + totalMarkersPrice
	
	var priceWithDiscount float32 = totalPrice - (totalPrice * float32(discount)/100)

	fmt.Println(priceWithDiscount)

}