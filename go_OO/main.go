package main

import (
	c "banco/contas"
	"fmt"
)

func main() {
	conta1 := c.ContaCorrente{"doglinha", 10, 123456, 1000}
	conta2 := c.ContaCorrente{"doglas", 20, 123654, 2000}

	fmt.Println(conta1, conta2)
	fmt.Println(conta2.Transferir(2000, &conta1))
	fmt.Println(conta1, conta2)
}
