//Faça um programa que preencha um vetor com 10 números inteiros. Calcule e mostre os números superiores a
//50 e suas respectivas posições. O programa deverá mostrar uma mensagem se não existir nenhum número
//nessa condição.

package main

import "fmt"

func main() {

	var numero[10]int
	var encontrou bool

	for i := 0; i < 10; i++ {
		fmt.Scan(&numero[i])
	}

	for indice, valor := range numero {
		if valor > 50 {			
			fmt.Println("Índice:", indice, ", Número:", valor)
			encontrou = true
		}
	}
	if !encontrou {
	fmt.Println("Nenhum número maior que 50.")
	}
}