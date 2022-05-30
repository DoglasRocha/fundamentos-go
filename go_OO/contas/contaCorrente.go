package contas

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return true
	}

	return false
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return true, c.saldo
	}

	return false, c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if c.Sacar(valorTransferencia) {
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
