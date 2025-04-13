//Um comerciante calcula o valor da venda, tendo em vista a tabela a seguir:
//Valor da Compra                 Valor da Venda
//Valor < R$ 10,00                Lucro de 70%
//R$ 10,00 <= Valor < R$ 30,00    Lucro de 50%
//R$ 30,00 <= Valor < R$ 50,00    Lucro de 40%
//Valor >= 50,00                  Lucro de 30%
//Escreva um programa que leia o valor da compra e imprima o valor da venda.
//Cuidado com valor inválido de compra!

package main

import "fmt"

func main() {

	var valorCompra, valorVenda float64

	fmt.Print("Digite o valor da comprar: ")
	fmt.Scan(&valorCompra)

	if valorCompra < 0 {
		fmt.Println("Valor inválido!")
	} else if valorCompra < 10 {
		valorVenda = valorCompra * 1.7
	} else if valorCompra >= 10 && valorCompra < 30 {
		valorVenda = valorCompra * 1.5
	} else if valorCompra >= 30 && valorCompra < 50 {
		valorVenda = valorCompra * 1.4
	} else if valorCompra >= 50 {
		valorVenda = valorCompra * 1.3
	}
	fmt.Printf("Valor da compra: R$ %.2f \n", valorCompra)
	fmt.Printf("Valor da venda: R$ %.2f \n", valorVenda)
}
