package main

import "fmt"

func main() {
	var nome string
	var comando int
	versao := 1.1

	fmt.Printf("Qual seu nome? ")
	fmt.Scanf("%s", &nome)

	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("Digite sua opção desejada:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	//fmt.Scanf("%d", &comando)
	// Scan infere a entrada baseado no tipo da variavel destino
	fmt.Scan(&comando)

}
