package main

import "fmt"

func main() {
	var a float64
	var b float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	pesoA := 3.5
	pesoB := 7.5

	result := (a*pesoA + b*pesoB) / (pesoA + pesoB)

	fmt.Printf("MEDIA = %.5f\n", result)

}
