//Crie um programa para determinar se um número inteiro A é divisível por outro
//número B. Os valores devem ser fornecidos pelo usuário.

package main

import "fmt"

func main() {

	var A, B int

	fmt.Println("Insira o valor de dois números inteiros:")
	fmt.Scan(&A, &B)

	if A%B == 0 {
		fmt.Println("O número é divisível.")
	} else {
		fmt.Println("O número não é divisível.")
	}
}
