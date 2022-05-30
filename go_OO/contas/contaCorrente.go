package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0

	if podeSacar {
		c.Saldo -= valorDoSaque
		return true
	}

	return false
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {

	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return true, c.Saldo
	}

	return false, c.Saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if c.Sacar(valorTransferencia) {
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
