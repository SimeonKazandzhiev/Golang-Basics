package main

import (
	"fmt"
	"math"
)

func main() {

	var inputNumber float32
	fmt.Scanln(&inputNumber)

	var result float32 = inputNumber * 180 / math.Pi

	fmt.Println(result)

}
