package main

import "fmt"

func main() {

    var A, B, C float32
    var discriminante float32
    
    fmt.Println("Insira os valores dos 3 coeficientes de uma equação quadrática:")
    fmt.Scan(&A, &B, &C)
      
    discriminante = (B * B) - (4 * A * C)
      
    fmt.Println("===================================")
    fmt.Printf("O VALOR DE DELTA E = %.2f \n", discriminante)
      
}