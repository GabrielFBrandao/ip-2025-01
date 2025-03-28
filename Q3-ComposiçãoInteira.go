package main

import "fmt"

func main() {

	var n1, n2, n3 int
	var concatenado, quadrado int

	fmt.Println("Digite 3 números inteiros maiores que 0 e menores que 10: ")
	fmt.Scan(&n1, &n2, &n3)

	if (n1 >= 10) || (n2 >= 10) || (n3 >= 10) || (n1 == 0) || (n2 == 0) || (n3 == 0) {
		fmt.Println("DÍGITO INVÁLIDO.")
	} else {
		n1 = n1 * 100
		n2 = n2 * 10
		concatenado = n1 + n2 + n3
		quadrado = concatenado * concatenado
		fmt.Println(concatenado, ",", quadrado)
	}
}
