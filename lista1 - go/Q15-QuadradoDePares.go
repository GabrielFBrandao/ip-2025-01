import "fmt"

func main() {

    var numero int
    var i int

    fmt.Println("Insira um valor inteiro, em que 5 < N < 2000: ")
    fmt.Scan(&numero)
      
    if numero < 5 || numero > 2000 {
        fmt.Println("NÚMERO INVÁLIDO.")
    } else {
        for i=0; i <= numero; i++ {
            if i % 2 == 0 && i != 0 {
                fmt.Printf("%d^2 = %d \n",i, i*i)
            }    
        }
    }
}