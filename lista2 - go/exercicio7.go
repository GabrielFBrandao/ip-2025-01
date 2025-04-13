//Escreva um programa que leia três valores inteiros distintos (assuma que o usuário
//digitará valores diferentes entre si) e os armazene nas variáveis A, B e C.
//Em seguida, descubra o menor valor e o armazene em uma variável denominada MENOR;
//o maior valor, coloque-o na variável MAIOR e o valor intermediário, na variável
//INTER. Imprima os valores em ordem crescente, ou seja, imprima as variáveis
//MENOR, INTER e MAIOR, nessa ordem.

package main

import "fmt"

func main() {

	var A, B, C int
	var MENOR, MAIOR, INTER int
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&A)
	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&B)
	fmt.Print("Digite o terceiro valor: ")
	fmt.Scan(&C)

	if A < B && A < C {
		MENOR = A
		if B < C {
			INTER = B
			MAIOR = C
			fmt.Println(MENOR, INTER, MAIOR)
		} else if B > C {
			INTER = C
			MAIOR = B
			fmt.Println(MENOR, INTER, MAIOR)
		}
	}
	if B < A && B < C {
		MENOR = B
		if A < C {
			INTER = A
			MAIOR = C
			fmt.Println(MENOR, INTER, MAIOR)
		} else if A > C {
			INTER = C
			MAIOR = A
			fmt.Println(MENOR, INTER, MAIOR)
		}
	}
	if C < B && C < A {
		MENOR = C
		if B < A {
			INTER = B
			MAIOR = A
			fmt.Println(MENOR, INTER, MAIOR)
		} else if B > A {
			INTER = A
			MAIOR = B
			fmt.Println(MENOR, INTER, MAIOR)
		}
	}
}
