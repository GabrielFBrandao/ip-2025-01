package main

import "fmt"

func main() {

    var x, y int
    var i int

    fmt.Print("Insira um número inteiro: ")
    fmt.Scan(&x)
    fmt.Print("Insira quantos números pares em sequência devem ser lidos: ")
    fmt.Scan(&y)
      
    fmt.Printf("%d  %d \n",x, y)
    fmt.Println("===============================")
      
    if x % 2 == 0 {
        for i=x; i <= (x + (2 * y) - 2); i++ {
            if i % 2 == 0 {
                fmt.Printf("%d  ", i)
            }
        }
    } else {
        fmt.Println("O PRIMEIRO NUMERO NAO É PAR.")
    }
}