package main

import "fmt"

func main(){

	var inputNum int;
	fmt.Scanln(&inputNum);

	if inputNum == 1 {
		fmt.Println("Monday")
	}else if inputNum == 2{
		fmt.Println("Tuesday")
	}else if inputNum == 3{
		fmt.Println("Wednesday")
	}else if inputNum == 4{
		fmt.Println("Thursday")
	}else if inputNum == 5{
		fmt.Println("Friday")
	}else if inputNum == 6{
		fmt.Println("Saturday")
	}else if inputNum == 7{
		fmt.Println("Sunday")
	}



}