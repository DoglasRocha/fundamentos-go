package main

import "fmt"
import "reflect"

func main() {

	// quando voce nao coloca valor na variavel, ela inicia zerada
	// atribuicao normal
	var nome string = "Doglas"
	var idade int = 18
	var versao float32 = 1.1

	// atribuicao por inferencia
	/* var nome = "Doglas" */

	// atribuicao curta
	/* nome := "Doglas" */

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	// pegando o tipo da variavel
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}
