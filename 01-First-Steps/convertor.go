package main

import "fmt"

func main() {

	var numToConvert float32
	var constant float32 = 2.54
	fmt.Scanln(&numToConvert)
	var result float32 = numToConvert * constant
	fmt.Printf("%.2f", result)
}
