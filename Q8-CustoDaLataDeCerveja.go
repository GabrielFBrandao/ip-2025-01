package main

import "fmt"

func main(){
    
    var r, a float32
    var area, valorTotal float32
    var valorPi float32
    
    fmt.Println("Insira o valor do raio e a altura da lata, respectivamente(m):")
    fmt.Scan(&r, &a)

    valorPi = 3.14159
    area = (2 * (valorPi * r * r)) + (2 * valorPi * r * a)
    valorTotal = area * 100

    fmt.Println("====================================")
    fmt.Printf("O VALOR DO CUSTO E = %.2f \n", valorTotal)
      
}