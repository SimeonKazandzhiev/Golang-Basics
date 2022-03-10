package main

import "fmt"

func main() {

	var architectName string
	fmt.Scanln(&architectName)
	var projectsNum int
	fmt.Scanln(&projectsNum)
	var totalProjectsTime int = projectsNum * 3

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.",architectName,totalProjectsTime,projectsNum)


}