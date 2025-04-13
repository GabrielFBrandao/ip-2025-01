/*Escrever um programa que leia o número de identificação, as 3 notas obtidas por 
um aluno nas 3 verificações e a média dos exercícios que fazem parte da avaliação. 
Calcular a média de aproveitamento do aluno, usando a fórmula:
Média Final = ( nota1 + nota2 ∗ 2 + nota3 ∗ 3 + média dos exercícios ) / 7
e o seu conceito, usando a tabela a seguir:
Média de Aproveitamento    Conceito
>= 9,0 e <= 10,0           A
>=7,5 e < 9,0              B
>=6,0 e < 7,5              C
>=4,0 e < 6,0              D
< 4,0                      E
O programa deve escrever o número do aluno, suas notas, a média dos exercícios, a 
média de aproveitamento, o conceito correspondente e a mensagem: APROVADO se o 
conceito for A, B ou C e REPROVADO, se o conceito for D ou E.*/

package main

import "fmt"

func main() {

	var identificacao int
	var n1, n2, n3, mediaEx, mediaFinal float64
	var conceito string

	fmt.Println("Insira o número de identificação do aluno: ")
	fmt.Scan(&identificacao)
	fmt.Println("Insira as 3 notas obtidas pelo aluno: ")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Println("Insira a média dos exercícios desse aluno: ")
	fmt.Scan(&mediaEx)

	mediaFinal = (n1 + (n2 * 2) + (n3 * 3) + mediaEx) / 7
	fmt.Printf("A média de aproveitamente foi: %.2f \n", mediaFinal)

	if mediaFinal >= 9 && mediaFinal <= 10 {
		conceito = "A"
	} else if mediaFinal >= 7.5 && mediaFinal < 9 {
		conceito = "B"
	} else if mediaFinal >= 6 && mediaFinal < 7.5 {
		conceito = "C"
	} else if mediaFinal >= 4 && mediaFinal < 6 {
		conceito = "D"
	} else if mediaFinal < 4 {
		conceito = "E"
	}
	
	switch conceito {
	case "A", "B", "C":
		fmt.Println("APROVADO")
	case "D", "E":
		fmt.Println("REPROVADO")
	}
	}