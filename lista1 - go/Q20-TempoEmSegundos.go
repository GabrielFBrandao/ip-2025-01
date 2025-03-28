package main

import "fmt"

func main() {

    var horas, minutos, segundos, soma int

    fmt.Println("Insira um número para as HORAS, MINUTOS e SEGUNDOS: ")
    fmt.Scan(&horas, &minutos, &segundos)
      
    soma = 0
    horas = horas * 3600
    minutos = minutos * 60
    soma = horas + minutos + segundos
      
    fmt.Println("=================================")
    fmt.Printf("O TEMPO EM SEGUNDOS E = %d \n", soma)
}