/*Desenvolver um programa que calcule o salário bruto e o salário líquido de um 
funcionário.
• Dados de Entrada: Matrícula do funcionário (int);
					Quantidade de horas-extras trabalhadas.
• Constantes: Salário Mínimo = R$ 788.00;
			  Valor da Hora-Extra = R$ 10.00.
Sabe-se:
• Salário hora-extra = horas-extras * Valor da Hora-Extra;
• Salário bruto = 3 * Salário Mínimo + Salário hora-extra;
• Desconto INSS = 12 % do salário bruto, se salário bruto for maior que R$ 1500,00;
• Desconto do Imposto de Renda = 20 % do Salário Bruto, se o mesmo for maior que R$ 2000,00;
• Salário líquido = salário bruto – deduções.*/

package main

import "fmt"

func main() {

	var matricula int
	var horasExtras, salarioHoraExtra, salarioBruto, salarioLiquido float64

fmt.Println("Qual a matrícula do funcionário?")
fmt.Scan(&matricula)
fmt.Println("Qual a quantidade de horas-extras trabalhadas?")
fmt.Scan(&horasExtras)

const salarioMinimo = 788
const valorHoraExtra = 10

salarioHoraExtra = horasExtras * valorHoraExtra
salarioBruto = (3 * salarioMinimo) + salarioHoraExtra

fmt.Printf("Salário Bruto é: %.2f \n", salarioBruto)

if salarioBruto > 2000 {
	salarioLiquido = salarioBruto * 0.88 * 0.8
} else if salarioBruto > 1500 && salarioBruto <= 2000 {
	salarioLiquido = salarioBruto * 0.88
} else {
	salarioLiquido = salarioBruto
}

fmt.Printf("Salário Líquido é: %.2f \n", salarioLiquido)
}