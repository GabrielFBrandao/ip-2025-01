package main

import (
	"fmt"
	"math")

func main() {
    
	var horas, aluguel float32
	
	// "_,err" foram declaradas e atribuídas ao valor que o usuário nos informará. Essa declaração com fmt.Scan necessita retornar 2 valores, por isso criamos 2 variáveis. O underline apenas retira a existência de uma delas, já que não será utilizada no código, por não ter utilidade.
	//Aqui, se o que for escrito pelo usuário não tiver nenhum erro, ou seja, err == nil, ele vai prosseguir com o código. Caso, err != nil, significa que existe algum erro, e a mensagem será posta na tela, com o erro evidenciado em seguida.
	fmt.Print("Insira a quantidade de horas que a charrete foi utilizada: ")
	_, err := fmt.Scan(&horas)
	if err != nil {
		fmt.Println("Erro na leitura do valor:", err)
		return
	}

	if horas <= 0 {
		fmt.Println("O número de horas deve ser positivo")
		return
	}

    //Aqui, entra a importação da biblioteca "math". A variável "horasInteiras" vai querer o valor inteiro das horas informadas pelo usuário. "math.Ceil" obrigatoriamente utiliza "float64".E como declaramos "horas" como float32, também terá de ser escrito e evidenciado para conversão.
    //Dessa forma, temos: int=valor inteiro das horas; math.Ceil=queremos o valor inteiro maior ou igual do que foi posto, ou seja, se as horas forem 3.4, será dado 4 como resultado; float64=a conversão necessária para os cálculos da biblioteca "math".
	horasInteiras := int(math.Ceil(float64(horas)))
	
	if horasInteiras%3 == 0 {
		aluguel = float32(horasInteiras/3) * 10
	} else {
		// Calcula os blocos de 3 horas e as horas adicionais
		blocos3horas := horasInteiras / 3
		horasExtras := horasInteiras % 3
		//Observe que "float32" sempre está aparecendo, indicando a conversão de float64, já que "blocos3horas" e "horasExtras" interagiu com "horasInteiras" para float32, relacionado com "aluguel".
		aluguel = float32(blocos3horas)*10 + float32(horasExtras)*5
	}

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", aluguel)
}