/*Criar um programa que leia a idade de uma pessoa e que mostre a sua classe eleitoral:
- Não-eleitor (abaixo de 16 anos);
- Eleitor obrigatório (entre 18 e 65 anos);
- Eleitor facultativo (entre 16 e 18 anos e maior de 65 anos).*/

package main

import "fmt"

func main() {

	var idade int 

	fmt.Print("Insira a idade da pessoa: ")
	fmt.Scan(&idade)

	if idade > 0 && idade < 16 {
		fmt.Println("Não-Eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Eleitor Obrigatório")
	} else if idade >= 16 && idade < 18 || idade > 65 {
		fmt.Println("Eleitor Facultativo")
	} else {
		return
	}
}