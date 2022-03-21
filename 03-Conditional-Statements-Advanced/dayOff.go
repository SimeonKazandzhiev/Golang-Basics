package main

import "fmt"

func main(){
	
	var day string
	fmt.Scanln(&day)

	if day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Thursday" || day == "Friday"{
		fmt.Println("Working day")
	}else if day == "Saturday" || day == "Sunday"{
		fmt.Println("Weekend")
	}else{
		fmt.Println("Error")
	}

}