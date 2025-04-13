/*Desenvolver um programa para calcular a conta de água para a SANEAGO. O custo da água
varia dependendo do tipo do consumidor - residencial, comercial ou industrial.
A regra para calcular a conta é:
• Residencial: R$ 5,00 de taxa mais R$ 0,05 por m3 gastos;
• Comercial: R$ 500,00 para os primeiros 80 m3 gastos mais R$ 0,25 por m3
gastos acima dos 80 m3;
• Industrial: R$ 800,00 para os primeiros 100 m3 gastos mas R$ 0,04 por m3
gastos acima dos 100 m3;
O programa deverá ler a conta do cliente, seu tipo (residencial, comercial e industrial)
e o seu consumo de água em metros cúbicos. Como resultado imprimir a conta do
cliente e o valor em real a ser pago pelo mesmo.*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var conta int
	var tipo string
	var consumo, valor float64

	fmt.Println("Qual a conta do cliente: ")
	fmt.Scan(&conta)
	fmt.Println("Qual o tipo do consumidor? (R-Residencial; C-Comercial; I-industrial)")
	fmt.Scan(&tipo)
	fmt.Println("Qual seu consumo em metros cúbicos(m3)?")
	fmt.Scan(&consumo)

	tipo = strings.ToUpper(tipo)

	switch tipo {
	case "R":
		if consumo <= 0 {
			fmt.Println("VALOR INCOERENTE.")
			return
		} else {
			valor = 5 + (0.05 * consumo)
		}
	case "C":
		if consumo <= 0 {
			fmt.Println("VALOR INCOERENTE.")
			return
		} else if consumo > 0 && consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (0.25 * (consumo - 80))
		}
	case "I":
		if consumo <= 0 {
			fmt.Println("VALOR INCOERENTE.")
			return
		} else if consumo > 0 && consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (0.04 * (consumo - 100))
		}
	default:
		fmt.Println("TIPO INDISPONÍVEL.")
		return
	}

	fmt.Printf("A conta do cliente é: %d \n", conta)
	fmt.Printf("O valor a ser pago será de R$%.2f,00 reais", valor)

}
