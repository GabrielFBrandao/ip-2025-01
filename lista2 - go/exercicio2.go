//Crie um programa que leia um valor inteiro
// e diga se ele é positivo, negativo ou nulo.

package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Insira um número inteiro: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("O número é positivo.")
	} else if numero == 0 {
		fmt.Println("O número é nulo.")
	} else if numero < 0 {
		fmt.Println("O número é negativo.")
	}
}
