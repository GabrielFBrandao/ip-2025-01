package main

import "fmt"

func main() {

	var salario, kW float32
	var custokW, consumo, desconto float32

	fmt.Println("Qual o valor do salário mínimo?(R$)")
	fmt.Scan(&salario)
	fmt.Println("Qual a quantidade de kW gasta na residência?")
	fmt.Scan(&kW)

	custokW = (0.7 / 100) * salario
	consumo = custokW * kW
	desconto = consumo * 0.9

	fmt.Println("===============================")
	fmt.Printf("Custo por kW: R$ %.2f \n", custokW)
	fmt.Printf("Custo do consumo: R$ %.2f \n", consumo)
	fmt.Printf("Custo com desconto: R$ %.2f \n", desconto)

}
