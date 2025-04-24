/*Escreva um programa que receba 10 números inteiros, armazene-os em um vetor e mostre:
a) os números pares digitados;
b) a soma dos números pares digitados;
c) os números ímpares digitados;
d) a quantidade de números ímpares digitados.*/

package main

import "fmt"

func main() {

	var numeros [10]int
	var vetorPares []int
	var vetorImpares []int
	var somaPares int
	var quantidadeImpares int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
	}

	for _, valor := range numeros {
		if valor % 2 == 0 {
			vetorPares = append(vetorPares, valor)
		}
	}

	somaPares = 0
	for _, valor := range numeros {
		if valor % 2 == 0 {
			somaPares += valor
		}
	}

	for _, valor := range numeros {
		if valor % 2 != 0 {
			vetorImpares = append(vetorImpares, valor)
		}
	}

	quantidadeImpares = 0
	for _, valor :=  range numeros {
		if valor % 2 != 0 {
			quantidadeImpares++
		}
	}
	fmt.Println("Os números pares digitados foram:", vetorPares)
	fmt.Println("A soma dos números pares digitados foram:", somaPares)
	fmt.Println("Os números ímpares digitados foram:", vetorImpares)
	fmt.Println("A quantidade de números ímpares digitados foram:", quantidadeImpares)
}