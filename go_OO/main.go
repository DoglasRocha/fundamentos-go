package main

import (
	"banco/cliente"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) bool
}

func main() {
	cliente1 := cliente.Titular{
		Nome:      "Doglas",
		CPF:       "999.999.999-99",
		Profissao: "Estudante",
	}

	conta1 := contas.ContaCorrente{
		Titular:       cliente1,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	conta2 := contas.ContaPoupanca{
		Titular:       cliente1,
		NumeroAgencia: 123,
		NumeroConta:   123654,
	}

	conta1.Depositar(5000)
	conta2.Depositar(5000)

	PagarBoleto(&conta1, 1000)
	PagarBoleto(&conta2, 500)

	fmt.Println(conta1)
	fmt.Println(conta2)
}
