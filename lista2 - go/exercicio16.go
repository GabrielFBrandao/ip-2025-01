/*Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do
segundo grau ( Ax2 + Bx + C = 0) e que calcule suas raízes. O programa deve mostrar,
quando possível, o valor das raízes calculadas e a classificação das mesmas:
“RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”.*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64
	var delta, r1, r2 float64

	fmt.Println("Insira os coeficientes A, B e C da equação do segundo grau: ")
	fmt.Scan(&a, &b, &c)

	delta = math.Pow(b, 2) - 4*a*c
	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS.")
		return
	} else if delta == 0 {
		fmt.Println("RAIZ ÚNICA.")
	} else {
		fmt.Println("RAÍZES DISTINTAS.")
	}

	r1 = (-b + math.Sqrt(delta)) / (2 * a)
	r2 = (-b - math.Sqrt(delta)) / (2 * a)

	fmt.Printf("Raiz 1: %.2f \n", r1)
	fmt.Printf("Raiz 2: %.2f \n", r2)
}
