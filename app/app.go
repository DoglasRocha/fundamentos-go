package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	// for "sem nada" em go é como se fosse um while (true)
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheco este comando")
		}
		fmt.Println()
	}
}

func exibeIntroducao() {
	var nome string
	versao := 1.2

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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	// parecido com for i, site in enumerate(sites) no python
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")

	} else {
		fmt.Println("Site:", site, "está com problemas. Error:", resp.StatusCode)
	}

	registraLog(site, resp.StatusCode == 200)
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") // simplesmente abre o arquivo
	// arquivo, err := ioutil.ReadFile("sites.txt") // abre o arquivo como um array de bytes, facil de trabalhar

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		sites = append(sites, strings.TrimSpace(linha))

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " +
		site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.log")

	if err != nil {
		fmt.Printf("Error reading logs:", err)
	}

	fmt.Println(string(arquivo))
}
