package main

import "fmt"

func main() {
	var number int
	var hour int
	var salary float64

	fmt.Scanln(&number)
	fmt.Scanln(&hour)
	fmt.Scanln(&salary)

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", salary*float64(hour))
}
