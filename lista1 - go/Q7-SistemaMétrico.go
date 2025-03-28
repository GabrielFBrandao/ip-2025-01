package main

import "fmt"

func main() {
 
    var fahrenheit, polegada float32
    var celsius, mm float32
    
    fmt.Println("Insira um valor em Fahrenheit e outro valor em polegadas: ")
    fmt.Scan(&fahrenheit, &polegada)

    celsius = ((5 * fahrenheit) - 160) / 9
    mm = polegada * 25.4
      
    fmt.Println("==========================================")
    fmt.Printf("O VALOR EM CELSIUS = %.2f ºC. \n", celsius)
    fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f mm. \n", mm)

}