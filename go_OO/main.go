package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaCorrente{"doglinha", 10, 123456, 4}

	fmt.Println(conta1)
}
