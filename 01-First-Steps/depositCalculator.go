package main

import "fmt"


func main() {

	var deposit float32
	fmt.Scanln(&deposit)
	var due int
	fmt.Scanln(&due)
	var anuallyPercent float32
	fmt.Scanln(&anuallyPercent)

	var interest float32 = deposit * anuallyPercent/100
	var interestPerMonth float32 = interest / 12 
	var totalSum float32 = deposit + float32(due) * interestPerMonth

	fmt.Println(totalSum)

}
