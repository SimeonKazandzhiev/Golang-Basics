package main

import "fmt"

func main () {



	var currentHour, currentMinutes int
	fmt.Scanln(&currentHour) 
	fmt.Scanln(&currentMinutes) 

	var currentTimeInMinutes int = currentHour * 60 + currentMinutes 
	var additionalTime int = currentTimeInMinutes + 15

	var finalHour int = additionalTime / 60  
	var finalMinutes int = additionalTime % 60 

	if (finalHour == 24) {
		finalHour = 0
	}

	fmt.Printf("%d:%02d", finalHour, finalMinutes)

}