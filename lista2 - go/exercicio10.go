//Escreva um programa que leia o destino de um passageiro e se a viagem inclui retorno
//(ou seja, se é só ida, ou se é ida e volta). Informe o preço da passagem de acordo
//com a tabela a seguir.
//Destino                  Ida         Ida e Volta
//1 - Região Norte         R$ 500,00   R$ 900,00
//2 - Região Nordeste      R$ 350,00   R$ 650,00
//3 - Região Centro-Oeste  R$ 350,00   R$ 600,00
//4 - Região Sul           R$ 300,00   R$ 550,00
//Sugestão: Considere “1” representando a Região Norte, “2” para Nordeste,
//“3” para Centro-Oeste e “4” para Sul. Para ver se a viagem inclui retorno,
//considere “1”- para sim (ida e volta) e “2” - para não (só ida).
//Cuidado com valores inválidos tanto para o destino quanto para o fato
//de incluir ou não o retorno.

package main

import "fmt"

func main() {

	var destino, retorno, valor int

	fmt.Println("Qual o seu destino? (1-Norte; 2-Nordeste; 3-Centro-Oeste; 4-Sul)")
	fmt.Scan(&destino)
	if destino < 1 && destino < 4 {
		fmt.Println("Destino Inválido!")
	}

	fmt.Println("A viagem inclui retorno? 1-Sim(ida e volta); 2-Não(só ida)")
	fmt.Scan(&retorno)
	if retorno < 1 && retorno > 2 {
		fmt.Println("Retorno Inválido!")
	}

	switch {
	case destino == 1 && retorno == 1:
		valor = 900
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 1 && retorno == 2:
		valor = 500
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 2 && retorno == 1:
		valor = 650
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 2 && retorno == 2:
		valor = 350
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 3 && retorno == 1:
		valor = 600
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 3 && retorno == 2:
		valor = 350
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 4 && retorno == 1:
		valor = 550
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	case destino == 4 && retorno == 2:
		valor = 300
		fmt.Printf("O valor a ser pago será de R$%d reais. \n", valor)
	}
}
