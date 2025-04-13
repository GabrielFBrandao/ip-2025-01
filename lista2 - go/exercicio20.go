/*Elabore um programa que calcule o valor a ser pago por um produto considerando o
preço normal de etiqueta e a escolha da condição de pagamento. Utilize os códigos
da tabela a seguir para saber qual a condição de pagamento escolhida e efetuar o
cálculo adequado.
Código    Condição de Pagamento
1         À vista, dinheiro ou cheque, 10% de desconto
2         À vista, cartão de crédito, 5% de desconto
3         Em 2 vezes, preço normal de etiqueta sem juros
4         Em 3 vezes, preço normal de etiqueta + 10% de juros*/

package main

import "fmt"

func main() {

	var preco float64
	var condicao int

	fmt.Println("Insira o valor do preço normal do produto:")
	fmt.Scan(&preco)
	fmt.Println("Qual será a condição de pagamento?(1-Dinheiro/Cheque; 2-Crédito; 3-Parcelado2x; 4-Parcelado3x)")
	fmt.Scan(&condicao)

	switch condicao {
	case 1:
		preco = preco * 0.9
	case 2:
		preco = preco * 0.95
	case 3:
		preco = preco
	case 4:
		preco = preco * 1.1
	default:
		fmt.Println("Valor Inválido.")
		return
	}

	fmt.Printf("O preço final do produto é: %.2f", preco)
}
