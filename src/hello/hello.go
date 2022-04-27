package main 

import "fmt"
import "reflect"

func main() {
	nome := "Enrico"
	idade := 19
	fmt.Println("Meu nome é", nome)
	fmt.Println("Eu tenho", idade, "anos")

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}