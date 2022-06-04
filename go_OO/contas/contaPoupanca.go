package contas

import (
	"banco/cliente"
)

type ContaPoupanca struct {
	Titular                              cliente.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return true
	}

	return false
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return true, c.saldo
	}

	return false, c.saldo
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if c.Sacar(valorTransferencia) {
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
