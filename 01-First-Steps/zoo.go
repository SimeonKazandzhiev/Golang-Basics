package main

import "fmt"

func main() {

	var dogsFoodPacks int
	fmt.Scanln(&dogsFoodPacks)

	var catsFoodPacks int
	fmt.Scanln(&catsFoodPacks)

	var totalSum float32 = (float32(dogsFoodPacks) * 2.5) + (float32(catsFoodPacks) * 4)

	fmt.Printf("%.1f lv.",totalSum)

}