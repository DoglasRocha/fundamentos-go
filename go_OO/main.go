package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// forma normal de usar struct
	conta1 := ContaCorrente{"doglinha", 10, 123456, 4}
	// forma passando o nome dos parametros
	conta2 := ContaCorrente{
		titular:       "doglao",
		numeroAgencia: 20,
		numeroConta:   1123654,
		saldo:         40}

	fmt.Println(conta1, conta2)

	// forma coisando cada parametro
	var conta3 *ContaCorrente   // cria um ponteiro para uma ContaCorrente
	conta3 = new(ContaCorrente) // liga o ponteiro a uma ContaCorrente
	conta3.titular = "doglas"

	fmt.Println(conta3)
}
