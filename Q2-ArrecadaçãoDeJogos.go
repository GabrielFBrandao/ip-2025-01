package main

import "fmt"

func main() {

	var casosTestes, i int
	var QuantPessoas [100]float32
	var CategPopular, CategGeral, CategArquibancada, CategCadeiras [100]float32
	var RendaTotal float32

	fmt.Println("Quantos casos testes de arrecadação você gostaria de analisar?")
	fmt.Scan(&casosTestes)

	fmt.Println("===============================================")
	for i = 0; i < casosTestes; i++ {
		fmt.Printf("DADOS RELACIONADOS À ARRECADAÇÃO DO JOGO Nº %d: \n", i+1)
		fmt.Print("Quantidade de pessoas: ")
		fmt.Scan(&QuantPessoas[i])
		fmt.Println(" ")
		fmt.Print("Porcentagem na Categoria Popular: ")
		fmt.Scan(&CategPopular[i])
		fmt.Println(" ")
		fmt.Print("Porcentagem na Categoria Geral: ")
		fmt.Scan(&CategGeral[i])
		fmt.Println(" ")
		fmt.Print("Porcentagem na Categoria Arquibancada: ")
		fmt.Scan(&CategArquibancada[i])
		fmt.Println(" ")
		fmt.Print("Porcentagem na Categoria Cadeiras: ")
		fmt.Scan(&CategCadeiras[i])
		fmt.Println("===================================================")
	}

	for i = 0; i < casosTestes; i++ {
		fmt.Println(QuantPessoas[i], CategPopular[i], CategGeral[i], CategArquibancada[i], CategCadeiras[i])
	}

	fmt.Println(" ")
	for i = 0; i < casosTestes; i++ {
		RendaTotal = 0
		RendaTotal = (QuantPessoas[i] * (CategPopular[i] / 100) * 1) + (QuantPessoas[i] * (CategGeral[i] / 100) * 5) + (QuantPessoas[i] * (CategArquibancada[i] / 100) * 10) + (QuantPessoas[i] * (CategCadeiras[i] / 100) * 20)
		fmt.Printf("A RENDA DO JOGO Nº %d foi de: %.2f \n", i+1, RendaTotal)
	}
}
