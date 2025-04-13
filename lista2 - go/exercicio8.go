//Faça um programa que indique se um número inteiro informado pelo usuário está
//compreendido entre 20 e 90 ou não. (20 e 90 não estão na faixa de valores).

package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&numero)

	if numero > 20 && numero < 90 {
		fmt.Println("O número está entre 20 e 90.")
	} else {
		fmt.Println("O número NÃO está entre 20 e 90.")
	}
}
