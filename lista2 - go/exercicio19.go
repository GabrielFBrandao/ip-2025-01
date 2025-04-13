/*Desenvolver um programa com as opções de calcular e imprimir o volume e a área da
superfície de um cone reto, de um cilindro ou de uma esfera. O programa deverá ler
a opção da figura desejada (1-cone / 2-cilindro / 3-esfera) e de acordo com a opção
escolhida calcular e escrever o volume e a área da superfície da figura pedida.*/

package main

import (
	"fmt"
	"math"
)

func main() {

	const PI = 3.14
	var figura int
	var raio, altura, volume, area float64

	fmt.Println("Insira a opção da figura deseja: (1-Cone; 2-Cilindro; 3-Esfera)")
	fmt.Scan(&figura)
	fmt.Println("Qual o raio e a altura dessa figura em metros, respectivamente?")
	fmt.Scan(&raio, &altura)

	switch figura {
	case 1:
		volume = (PI * math.Pow(area, 2) * altura) / 3
		area = PI * raio * math.Sqrt(math.Pow(raio, 2)+math.Pow(altura, 2))
	case 2:
		volume = PI * math.Pow(raio, 2) * altura
		area = 2 * PI * raio * altura
	case 3:
		volume = (4 * PI * math.Pow(raio, 3)) / 3
		area = 4 * PI * math.Pow(raio, 2)
	}

	fmt.Printf("O volume da figura é: %.2f m³\n", volume)
	fmt.Printf("A área da figura é: %.2f m²\n", area)

}
