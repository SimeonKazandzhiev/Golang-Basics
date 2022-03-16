package main

import "fmt"

func main () {

	var budget float32
	fmt.Scanln(&budget) 

	var countVideoCards, countCPUs, countRam int
	fmt.Scanln(&countVideoCards) 
	fmt.Scanln(&countCPUs)
	fmt.Scanln(&countRam) 
	var sumVideoCards float32 = float32(countVideoCards) * 250.0 
	var sumCPUs float32 = float32(countCPUs) * (0.35 * sumVideoCards)
	var sumRam float32 = float32(countRam) * (0.10 * sumVideoCards) 

	var totalSum float32 = sumVideoCards + sumCPUs + sumRam 

	if countVideoCards > countCPUs {
		totalSum = totalSum - 0.15 * totalSum 
	}

	
	if budget >= totalSum {
		var leftMoney float32 = budget - totalSum
		fmt.Printf("You have %.2f leva left!", leftMoney)
	} else {
		var needSum float32 = totalSum - budget
		fmt.Printf("Not enough money! You need %.2f leva more!", needSum)
	}

}