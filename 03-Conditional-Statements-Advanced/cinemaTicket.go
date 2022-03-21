package main

import "fmt"

func main() {

	var day string
	fmt.Scanln(&day)

	switch day {
	case "Monday":
		fmt.Println(12)
	case "Tuesday":
		fmt.Println(12)
	case "Wednesday":
		fmt.Println(14)
	case "Thursday":
		fmt.Println(14)
	case "Friday":
		fmt.Println(12)
	case "Saturday":
		fmt.Println(16)
	case "Sunday":
		fmt.Println(16)

	}

}
