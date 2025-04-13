//Elabore um programa que leia dois números inteiros e efetue a adição dos mesmos.
//Caso o valor somado seja maior do que 20, o resultado a ser apresentado será a
//soma mencionada adicionada de 8. Caso o valor somado seja menor ou igual a 20,
//deve-se subtrair 5 do mesmo para apresentá-lo em seguida.

package main

import "fmt"

func main() {

	var n1, n2, soma int

	fmt.Println("Insira o valor de dois números inteiros: ")
	fmt.Scan(&n1, &n2)

	soma = n1 + n2

	if soma > 20 {
		soma = soma + 8
	} else if soma <= 20 {
		soma = soma - 5
	}

	fmt.Println("O valor da nova soma é: ", soma)
}
