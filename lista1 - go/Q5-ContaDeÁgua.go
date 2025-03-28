package main

import (
	"fmt"
	"strings"
)

func main() {

	var conta int
	var consumo float32
	var consumidor string
	var valorConta float32

	fmt.Print("Valor da conta do cliente: ")
	fmt.Scan(&conta)
	fmt.Print("Consumo de água por metro cúbico: ")
	fmt.Scan(&consumo)
	fmt.Print("O tipo de consumidor(C - COMERCIAL, I - INDUSTRIAL ou R - RESIDENCIAL): ")
	fmt.Scan(&consumidor)
	fmt.Println(conta, consumo, " ", consumidor)

	consumidor = strings.ToUpper(consumidor)

	switch consumidor {
	case "R":
		valorConta = 5 + (0.05 * consumo)
	case "C":
		valorConta = 500 + (0.25 * (consumo - 80))
	case "I":
		valorConta = 800 + (0.04 * (consumo - 100))
	}

	fmt.Println("===================================")
	fmt.Printf("CONTA = %d \n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f \n", valorConta)
}
