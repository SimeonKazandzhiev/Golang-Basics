package main

import "fmt"

func main (){

	var age float32
	fmt.Scanln(&age)
	var gender string
	fmt.Scanln(&gender)

	if age >= 16 && gender == "m"{
		fmt.Println("Mr.")
	}else if age < 16 && gender == "m"{
		fmt.Println("Master")
	}

	if age >= 16 && gender == "f"{
		fmt.Println("Ms.")
	}else if age < 16 && gender == "f"{
		fmt.Println("Miss")
	}
}