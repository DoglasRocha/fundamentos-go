package main

import (
	"fmt"
	"go_OO/contas"
)

func main() {
	conta1 := contas.ContaCorrente{"doglinha", 10, 123456, 1000}
	conta2 := contas.ContaCorrente{"doglas", 20, 123654, 2000}

	fmt.Println(conta1, conta2)
	fmt.Println(conta1.Transferir(2000, &conta2))
	fmt.Println(conta1, conta2)
}
