//Escreva um programa que leia um número inteiro e diga se ele é ou não
//um número divisível por 5.

package main

import "fmt"

func main() {

	var numero int

	fmt.Println("Insira um número inteiro:")
	fmt.Scan(&numero)

	if numero%5 == 0 {
		fmt.Println("O número é divisível.")
	} else {
		fmt.Println("O número não é divisível.")
	}
}
