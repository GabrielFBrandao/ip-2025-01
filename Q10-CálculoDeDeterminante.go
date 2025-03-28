package main

import "fmt"

func main() {

    var n1, n2, n3, n4 float32
    var determinante float32
    
    fmt.Println("Insira os 4 valores de uma matriz 2x2:")
    fmt.Scan(&n1, &n2, &n3, &n4)

    determinante = (n1 * n4) - (n2 * n3)

      fmt.Println(" ")
      fmt.Println("REPRESENTAÇÃO DA MATRIZ 2X2:")
      fmt.Println(n1, n2)
      fmt.Println(n3, n4)
      fmt.Println("=====================================")
      fmt.Printf("O VALOR DO DETERMINANTE E = %f \n", determinante)
      
}