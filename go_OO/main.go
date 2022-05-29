package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func main() {
	conta1 := ContaCorrente{"doglinha", 10, 123456, 1000}

	fmt.Println(conta1)

	fmt.Println(conta1.Sacar(399))

	fmt.Println(conta1)
}
