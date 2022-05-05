package main 

import (
	"fmt"
	"os"
	"net/http"
)

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
	sites := []string{"https://www.youtube.com", "https://www.github.com"}
	
	for i, site := range sites {
		fmt.Println("Testando o site", i,":", site)
		testeSite(site)
	}
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
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println(site, "Foi carregado com sucesso")
		return;
	} 

	fmt.Println(site, "Está fora do ar :(")
		
}