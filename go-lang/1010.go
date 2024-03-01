package main

import (
	"fmt"
)

func main() {
	var (
		cod   int
		qtd   int
		valor float64

		cod2   int
		qtd2   int
		valor2 float64
	)

	fmt.Scanf("%d %d %f\n", &cod, &qtd, &valor)
	fmt.Scanf("%d %d %f\n", &cod2, &qtd2, &valor2)

	total := valor*float64(qtd) + valor2*float64(qtd2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
