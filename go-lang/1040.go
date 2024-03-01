package main

import "fmt"

func main() {
	var (
		a, b, c, d float64
	)
	fmt.Scanf("%f %f %f %f\n", &a, &b, &c, &d)

	media := (a*2 + b*3 + c*4 + d*1) / 10

	fmt.Printf("Media: %.1f\n", media)

	if media >= 7 {
		fmt.Println("Aluno aprovado.")
	}
	if media < 5 {
		fmt.Println("Aluno reprovado.")
	}
	if media >= 5 && media < 7 {
		fmt.Println("Aluno em exame.")
		var exame float64
		fmt.Scanf("%f", &exame)
		fmt.Printf("Nota do exame: %.1f\n", exame)
		media = (media + exame) / 2
		if media >= 5 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", media)
	}
}
