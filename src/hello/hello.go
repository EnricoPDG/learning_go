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

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
	
		switch comando {
		case 1:
			inicarMonitoranmento()	
		case 2: 
			fmt.Println("Exibindo log")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse programa")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Enrico"
	var versao float32 = 1.0
	fmt.Println("Olá", nome)
	fmt.Println("A versão desse programa é", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir log")
	fmt.Println("0- Sair do programa")
}

func inicarMonitoranmento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	
	for i := 0; i < monitoramento; i++{	
		for i, site := range sites {
			fmt.Println("Testando o site", i,":", site)
			testeSite(site)		
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

 func leComando() int {
	var comandoLido int


	fmt.Println("Digite o comando: ")

	//& -> Passa o endereço de memória
	fmt.Scan(&comandoLido)

	fmt.Println("Você digitou o comando:", comandoLido)
	fmt.Println("Que está alocado no endereço de memória:", &comandoLido)

	return comandoLido
} 

func testeSite(site string) {
	resp, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "Foi carregado com sucesso")
		registraLog(site, true)
		
	}else {
		registraLog(site, false)
		fmt.Println(site, "Está fora do ar :(")		
	}
	
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu este erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu este erro:", err)
	}

	fmt.Println(string(arquivo))

}