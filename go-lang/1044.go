package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if b%a == 0 || a%b == 0 {
		fmt.Println("Sao Multiplos")
		return
	}

	fmt.Println("Nao sao Multiplos")
}
