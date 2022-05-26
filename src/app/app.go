package main

import "fmt"
import "os"

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo logs...")

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheco este comando")
	}
}

func exibeIntroducao() {
	var nome string
	versao := 1.1

	fmt.Printf("Qual seu nome? ")
	fmt.Scanf("%s", &nome)

	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("Digite sua opção desejada:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
