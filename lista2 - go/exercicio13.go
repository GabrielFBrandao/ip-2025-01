//Escreva um programa que receba um número inteiro positivo de 3 casas e imprima
//o algarismo da casa das dezenas. Não se esqueça de testar para ver se
//o número informado tem realmente 3 casas.

package main

import "fmt"

var numero int

func main() {

	fmt.Print("Insira um número inteiro com TRÊS casas: ")
	fmt.Scan(&numero)

	if numero < 100 && numero > 999 {
		fmt.Println("Número Inválido.")
	} else {
		numero = numero / 10
		numero = numero % 10
	}

	fmt.Printf("O algarismo da casa das dezenas é: %d \n", numero)
}
