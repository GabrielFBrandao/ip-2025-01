package main

import "fmt"

func main() {

	var quantidade, i int
	var nF [100]float32
	var nC [100]float32

	fmt.Println("Quantos valores de temperatura serão convertidos?")
	fmt.Scan(&quantidade)

	fmt.Println("Insira os respectivos valores de temperaturas em Fahrenheit(Fº):")
	for i = 0; i < quantidade; i++ {
		fmt.Scan(&nF[i])
	}

	fmt.Println("================================")

	for i = 0; i < quantidade; i++ {
		nC[i] = (5 * (nF[i] - 32)) / 9
	}

	for i = 0; i < quantidade; i++ {
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS. \n", nF[i], nC[i])
	}
}
