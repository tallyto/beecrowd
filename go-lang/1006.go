package main

import "fmt"

func main() {
	var nota1 float64
	var nota2 float64
	var nota3 float64

	var peso1 = 2.0
	var peso2 = 3.0
	var peso3 = 5.0

	fmt.Scanln(&nota1)
	fmt.Scanln(&nota2)
	fmt.Scanln(&nota3)

	var media = (nota1*peso1 + nota2*peso2 + nota3*peso3) / (peso1 + peso2 + peso3)

	fmt.Printf("MEDIA = %.1f", media)
}
