package main

import "fmt"

func main(){

	var flowerKind string
	fmt.Scanln(&flowerKind)
	var flowersCount float32
	fmt.Scanln(&flowersCount)
	var budget float32
	fmt.Scanln(&budget)

	var finalPrice float32

	if flowerKind == "Roses"{
		if flowersCount > 80{
			finalPrice = flowersCount * 5 * 0.9
		}else{
			finalPrice = flowersCount * 5
		}	
	}else if flowerKind == "Dahlias"{
		if flowersCount > 90 {
			finalPrice = flowersCount * 3.80 * 0.85
		}else{
			finalPrice = flowersCount * 3.80
		}
	}else if flowerKind == "Tulips"{
		if flowersCount > 80{
			finalPrice = flowersCount * 2.80 * 0.85
		}else{
			finalPrice = flowersCount * 2.80
		}
	}else if flowerKind == "Narcissus"{
		if flowersCount < 120{
			finalPrice = flowersCount * 3 * 1.15
		}else{
			finalPrice = flowersCount * 3
		}
	}else if flowerKind == "Gladiolus"{
		if flowersCount < 80 {
			finalPrice = flowersCount * 2.50 * 1.20
		}else{
			finalPrice = flowersCount * 2.50
		}
	}
	
	if finalPrice > budget {
		fmt.Printf("Not enough money, you need %.2f leva more.", finalPrice - budget)
	}else{
		fmt.Printf("Hey, you have a great garden with %.0f %s and %.2f leva left.",flowersCount,flowerKind,budget-finalPrice)
	}

}