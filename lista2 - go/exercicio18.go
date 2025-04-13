/*Uma locadora tem as seguintes regras para aluguel de DVDs:
- Às segundas, terças e quintas (2, 3 e 5) : desconto de 40% em relação ao preço normal;
- Às quartas , sextas, sábados e domingos (4,6 ,7 e 1): preço normal;
- Aluguel de DVDs comuns: preço normal;
- Aluguel de lançamentos: acréscimo de 15% em relação ao preço normal.
Desenvolver um programa para ler o preço normal do DVD alugado (em R$) e sua categoria
(comum ou lançamento). Calcular e imprimir o preço final que será pago pela locação do DVD.*/

package main

import "fmt"

var preco float64
var categoria string
var dia int

func main() {

	fmt.Println("Insira o preço normal do DVD: ")
	fmt.Scan(&preco)

	fmt.Println("O DVD alugado era comum ou dede lançamento? (C/L)")
	fmt.Scan(&categoria)
	if categoria == "C" {
		preco = preco
	} else if categoria == "L" {
		preco = preco * 1.15
	}

	fmt.Println("Qual o dia de locação? (2-Seg; 3-Ter; 4-Qua; 5-Qui; 6-Sex; 7-Sab; 8-Dom)")
	fmt.Scan(&dia)

	switch dia {
	case 2, 3, 5:
		preco = preco * 0.6
	case 4, 6, 7, 1:
		preco = preco
	}

	fmt.Printf("O preço final de locação desse DVD é de R$%.2f,00 reais", preco)
}
