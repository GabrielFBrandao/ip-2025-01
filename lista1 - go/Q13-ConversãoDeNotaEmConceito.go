package main

import "fmt"

func main() {
    
    var nota float32
    var conceito string
   
    fmt.Print("Insira uma nota de 0 a 10 em valor real: ")
    fmt.Scan(&nota)
      
    if nota <= 10 && nota >= 9 {
        conceito = "A"
    } else if nota < 9 && nota >= 7.5 {
        conceito = "B"
    } else if nota < 7.5 && nota >= 6 {
        conceito = "C"
    } else if nota < 6 && nota >= 0 {
        conceito = "D"
    }
                     
    fmt.Printf("NOTA = %.2f \nCONCEITO = %s \n",nota, conceito)
}