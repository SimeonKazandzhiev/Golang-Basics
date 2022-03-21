package main

import "fmt"

func main() {

	var ticket string
	fmt.Scanln(&ticket)
	var rows int
	fmt.Scanln(&rows)
	var cols int
	fmt.Scanln(&cols)

	var finalPrice float32

	switch ticket {
	case "Premiere":
		finalPrice = float32(rows) * float32(cols) * 12.00
	case "Normal":
		finalPrice = float32(rows) * float32(cols) * 7.50
	case "Discount":
		finalPrice = float32(rows) * float32(cols) * 5.00

	}

	fmt.Printf("%.2f leva",finalPrice)

}
