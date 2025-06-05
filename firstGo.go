package main // Principal pacote, código Go deve começar com o pacote main!
import (
	"fmt"
	"net/http"
	"os"
)

func introduction() {
	fmt.Println("Bem vindo ao programa de monitoramento de sites!, Por favor, digite seu nome para continuar:")
	var nome string
	fmt.Scan(&nome)
	var version = 1.1
	fmt.Println("Olá, Sr(a)", nome, "Seja bem vindo ao programa de monitoramento de sites!")
	fmt.Println("Versão do programa:", float32(version))
}

func serverWeb() { // Função para monitorar o servidor web, onde o usuário informa o site a ser monitorado
	fmt.Println(("Digita o site a ser monitorado (exemplo: https://www.google.com):"))
	site := []string{"https://www.google.com", "https://instagram.com.br", "https://facebook.com"} // Slice de strings para armazenar o site a ser monitorado - É um array dinâmico, que pode crescer ou diminuir conforme necessário
	for _, site := range site {                                                                    // Loop para percorrer o slice de sites e imprimir cada um deles
		testSite(site) // Chama a função testSite para verificar se o site está ativo
	}
}
func testSite(site string) { // Função para testar se o site está ativo, onde o usuário informa o site a ser monitorado
	fmt.Println("Monitorando o site:", site)
	resp, _ := http.Get(site) // Realiza uma requisição HTTP GET para o site informado pelo usuário. Ignorando o erro retornado com "_"
	if resp.StatusCode == 200 {
		fmt.Println("Site carregado com sucesso! Status: ", resp.StatusCode)
		fmt.Println("")
	} else {
		fmt.Println("O site não foi carregado, verifique sua conexão ou se o site está ativo!")
		fmt.Println("O site retornou o código de status:", resp.StatusCode)
		os.Exit(-1)
		fmt.Println("")
	}

}

func menu() {

	fmt.Println("Por favor, selecione o tipo de servidor que deseja monitorar:")
	fmt.Println("1 - Servidor Web")
	fmt.Println("2 - Logs")
	fmt.Println("3 - Sair do Programa")

}
func lercomando() int {
	var comandolido int
	fmt.Scan(&comandolido)
	return comandolido
}

func main() {

	introduction()
	for {
		menu()
		comando := lercomando()
		switch comando { // Switch case para verificar o comando digitado pelo usuário | Pode utilzar o if, porémo switch case é mais legível e organizado

		case 1:
			fmt.Println("Você escolheu monitorar um site Web!")
			serverWeb()
			// break - Não é necessário, pois o switch case já finaliza após o comando ser executado
		case 2:
			fmt.Println("Você escolheu visualizar os Logs...")
		case 3:
			fmt.Println("Obrigado por usar o programa de monitoramento de sites!")
			fmt.Print("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido! Por favor, selecione uma opção válida.")
			os.Exit(-1)
		}
	}
}

// func Sites() {
// 	sites := []string{} // Slice de strings para armazenar os sites a serem monitorados - É um array dinâmico, que pode crescer ou diminuir conforme necessário
// 	fmt.Println("Digite os sites a serem monitorados (digite 'sair' para finalizar):")
// 	for {
// 		var site string
// 		fmt.Scan(&site)
// 		if site == "sair" {
// 			break // Sai do loop se o usuário digitar "sair"
// 		}
// 		sites = append(sites, site) // Adiciona o site ao slice
// 		fmt.Println("Site adicionado:", site)
// 	}
// }
