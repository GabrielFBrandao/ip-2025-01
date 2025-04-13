//Desenvolva um programa para calcular e imprimir o preço final de um carro.
//O valor do preço inicial de fábrica é fornecido pelo usuário.
//O carro pode ter as seguintes opções:
//a) Ar condicionado (R$ 1750,00)
//b) Pintura metálica (R$ 800,00)
//c) Vidro elétrico (R$ 1200,00)
//d) Direção hidráulica (R$ 2000,00)

package main

import "fmt"

func main() {

	var carro int
	var arCondicionado, pinturaMetalica, vidroEletrico, direcaoHidraulica string

	fmt.Print("Insira o valor inicial do carro: ")
	fmt.Scan(&carro)

	if carro < 0 {
		fmt.Println("Valor Inválido.")
		return
	}

	fmt.Print("O carro tem ar condicionado? (S/N): ")
	fmt.Scan(&arCondicionado)
	if arCondicionado == "S" {
		carro += 1750
	} else if arCondicionado == "N" {
		carro += 0
	} else {
		fmt.Println("Opção Inválida.")
		return
	}

	fmt.Print("O carro tem pintura metálica? (S/N): ")
	fmt.Scan(&pinturaMetalica)
	if pinturaMetalica == "S" {
		carro += 800
	} else if pinturaMetalica == "N" {
		carro += 0
	} else {
		fmt.Println("Opção Inválida.")
		return
	}

	fmt.Print("O carro tem vidro elétrico? (S/N): ")
	fmt.Scan(&vidroEletrico)
	if vidroEletrico == "S" {
		carro += 1200
	} else if vidroEletrico == "N" {
		carro += 0
	} else {
		fmt.Println("Opção Inválida.")
		return
	}

	fmt.Print("O carro tem direção hidráulica? (S/N): ")
	fmt.Scan(&direcaoHidraulica)
	if direcaoHidraulica == "S" {
		carro += 2000
	} else if direcaoHidraulica == "N" {
		carro += 0
	} else {
		fmt.Println("Opção Inválida.")
		return
	}

	fmt.Println("O preço final do carro é: R$", carro, "reais.")
}
