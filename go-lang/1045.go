package main

import "fmt"

func main() {
	var a, b, c, aux float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	if b > a && b > c {
		aux = b
		b = c
		c = a
		a = aux
	} else {
		if c > a && c > b {
			aux = c
			c = b
			b = a
			a = aux
		}
	}

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {

		if a*a == b*b+c*c {
			fmt.Println("TRIANGULO RETANGULO")
		} else if a*a < b*b+c*c {
			fmt.Println("TRIANGULO ACUTANGULO")
			if a == b && b == c && c == a {
				fmt.Println("TRIANGULO EQUILATERO")
			} else if a == b || a == c || c == b {
				fmt.Println("TRIANGULO ISOSCELES")
			}
		} else if a*a > (b*b + c*c) {
			fmt.Println("TRIANGULO OBTUSANGULO")
			if a == b && b == c && c == a {
				fmt.Println("TRIANGULO EQUILATERO")
			} else if a == b || a == c || c == b {
				fmt.Println("TRIANGULO ISOSCELES")
			}
		}

	}

}
