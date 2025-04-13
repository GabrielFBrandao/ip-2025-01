//Escreva um programa que leia um valor inteiro
// e diga se o número informado é par ou ímpar.

package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Insira um número inteiro:")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("O número é par!")
	} else {
		fmt.Println("O número é ímpar!")
	}
}
