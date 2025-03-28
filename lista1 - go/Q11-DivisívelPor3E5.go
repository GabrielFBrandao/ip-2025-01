package main 

import "fmt"

func main() {

    var numero int

    fmt.Println("Insira um número inteiro: ")
    fmt.Scan(&numero)

    if (numero % 3 == 0) && (numero % 5 == 0) {
        fmt.Println("O NÚMERO É DIVISÍVEL!")
    } else {
        fmt.Println("O NÚMERO NÃO É DIVISÍVEL!")
    }
  
}