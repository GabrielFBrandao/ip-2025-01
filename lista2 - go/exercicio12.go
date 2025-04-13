//Construa um programa que receba a idade de uma pessoa e classifique-a seguindo o
//critério apresentado a seguir. Considere que a idade é um valor inteiro
//e que será informada de forma válida.
//Idade             Classificação
//0 a 2 anos        Recém-nascido
//3 a 11 anos       Criança
//12 a 19 anos      Adolescente
//20 a 55 anos      Adulto
//Acima de 55 anos  Idoso

package main

import "fmt"

func main() {

	var idade int

	fmt.Println("Qual a idade da pessoa?")
	fmt.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		fmt.Println("Classificado como Recém-Nascido.")
	} else if idade >= 3 && idade <= 11 {
		fmt.Println("Classificado como Criança.")
	} else if idade >= 12 && idade <= 19 {
		fmt.Println("Classificado como Adolescente.")
	} else if idade >= 20 && idade <= 55 {
		fmt.Println("Classificado como Adulto.")
	} else if idade > 55 {
		fmt.Println("Classificado como Idoso.")
	}
}
