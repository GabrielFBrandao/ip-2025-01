//Escreva um programa que receba o valor de x, calcule e imprima o valor de f(x),
//dado que: f(x) = 8 / (2-x)

package main

import "fmt"

func main() {

	var valor, fx float64

	fmt.Print("Insira um valor para x: ")
	fmt.Scan(&valor)

	fx = 8 / (2 - valor)

	fmt.Printf("O valor da função f(x) é: %.2f", fx)
}
