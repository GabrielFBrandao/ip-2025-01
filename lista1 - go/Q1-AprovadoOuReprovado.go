package main

import "fmt"

func main() {

	var N [3]int
	var media, soma float32
	var i int

	for i = 0; i < 3; i++ {
		fmt.Printf("Qual a %d° nota do aluno? \n", i+1)
		fmt.Scan(&N[i])
	}

	soma = 0
	for i = 0; i < 3; i++ {
		soma += float32(N[i])
	}

	media = 0
	media = soma / 3

	fmt.Printf("MÉDIA: %.2f \n", media)
	if media >= 6 {
		fmt.Println("APROVADO.")
	} else {
		fmt.Println("REPROVADO.")
	}
}
