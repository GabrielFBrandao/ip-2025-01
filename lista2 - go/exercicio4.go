//Faça um programa que leia um número do tipo float e imprima sua raiz quadrada
//caso o mesmo seja positivo ou nulo. Caso ele seja negativo, mostre o seu quadrado.

package main

import (
	"fmt"
	"math"
)

func main() {

	var numero float64

	fmt.Println("Insira um número real: ")
	fmt.Scan(&numero)

	if numero >= 0 {
		fmt.Printf("Sua raiz quadrada é: %.2f ", math.Sqrt(numero))
	} else {
		fmt.Printf("Seu quadrado é: %.2f", math.Pow(numero, 2))
	}
}
