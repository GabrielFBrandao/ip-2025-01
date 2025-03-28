package main

import ("fmt"
        "math")
        
func main() {

    var altura, aresta float64
    var volume float64

    fmt.Print("Insira o valor da altura da pirâmide(m): ")
    fmt.Scan(&altura)
    fmt.Print("Insira o valor da aresta da base da pirâmide(m): ")
    fmt.Scan(&aresta)
      
    fmt.Println(altura)
    fmt.Println(aresta)

    volume = (math.Pow(aresta, 2) * math.Sqrt(3) * altura) / 2
      
    fmt.Println("=======================================")
    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS. \n", volume)
}