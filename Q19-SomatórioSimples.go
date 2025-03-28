package main

import "fmt"

func main() {

    var numero, i int
    var soma float32

    fmt.Print("Insira um número inteiro positivo maior que um: ")
    fmt.Scan(&numero)
    
    if numero <= 1 {
        fmt.Println("Número inválido.")
        return
    } else {
        soma = 0
        for i=1; i <= numero; i++ {
            soma += 1/ float32(i)
        }
    }
      
    fmt.Printf(" %.6f \n", soma)
}