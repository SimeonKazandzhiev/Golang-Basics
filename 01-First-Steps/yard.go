package main

import "fmt"

func main() {

	var landToBeGreen float32
	fmt.Scan(&landToBeGreen)

	var totalPrice float32 = landToBeGreen * 7.61
	var discount float32 = totalPrice * 0.18
	var finalPrice float32 = totalPrice - discount
	
	fmt.Printf("The final price is: %.2f lv.", finalPrice)
	fmt.Println()
	fmt.Printf("The discount is: %.2f lv.",discount)

}