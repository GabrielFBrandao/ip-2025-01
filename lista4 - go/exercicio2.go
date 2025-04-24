/*Faça um programa que preencha um vetor com 10 números inteiros e um segundo vetor com 5 elementos
inteiros, ambos informados pelo usuário. Calcule e mostre dois vetores resultantes. O primeiro vetor resultante
será composto pelos números pares do primeiro vetor somado a todos os elementos do segundo vetor. O
segundo será composto pelos números ímpares do primeiro vetor somado a todos os elementos do segundo
vetor.
Exemplo:
Primeiro vetor: [ 4 7 5 8 2 15 9 6 10 11 ]
Segundo vetor: [ 3 4 5 8 2 ]
Primeiro vetor resultante: [ 26 30 24 . . . ], onde 26 = 4+3+4+5+8+2, 30 = 8+3+4+5+8+2, ...
Segundo vetor resultante: [ 29 27 37 . . .], onde 29 =7+3+4+5+8+2, 27 = 5+3+4+5+8+2, ...*/

package main

import "fmt"

func main() {

	var vetor1 [10]int
	var vetor2 [5]int
	var vetorResultantePares []int
	var vetorResultanteImpares []int

	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor2[i])
	}

	for _, valor := range vetor1 {
		soma := 0
		for _, v := range vetor2 {
			soma += v 
		}
		if valor % 2 == 0 {
			vetorResultantePares = append(vetorResultantePares, valor+soma)
		} else {
			vetorResultanteImpares = append(vetorResultanteImpares, valor+soma)
		}
	}

	fmt.Println("Vetor Resultante dos Pares:", vetorResultantePares)
	fmt.Println("Vetor Resultante dos Ímpares:", vetorResultanteImpares)
}