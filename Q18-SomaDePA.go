package main

import "fmt"

func main() {

    var primeiro, razao, quantidade int
    var soma, i int

    fmt.Print("Insira o primeiro número da PA: ")
    fmt.Scan(&primeiro)
    fmt.Print("Insira a razão da PA: ")
    fmt.Scan(&razao)
    fmt.Print("Insira o número de elementos da PA: ")
    fmt.Scan(&quantidade)
      
    fmt.Println(primeiro, " ", razao, " ", quantidade)
      
    soma = 0
    for i=0; i < quantidade; i++ {
        soma = soma + (primeiro + ((i - 1) * razao))
    }
      
    fmt.Println(soma)
}