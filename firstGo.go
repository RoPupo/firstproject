package main // Principal pacote, código Go deve começar com o pacote main!
import (
	"fmt"
)

func introduction() {
	nome := "Rodrigo"
	var version = 1.1
	fmt.Println("Olá, Sr", nome, "Seja bem vindo ao programa de monitoramento de servidores!")
	fmt.Println("Versão do programa:", float32(version))
}

func main() {

	introduction()

	fmt.Println("Por favor, selecione o tipo de servidor que deseja monitorar:")
	fmt.Println("1 - Servidor Web")
	fmt.Println("2 - Servidor de Banco de Dados")
	fmt.Println("3 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	switch comando { // Switch case para verificar o comando digitado pelo usuário
	// Pode utilzar o if, porémo switch case é mais legível e organizado
	case 1:
		fmt.Println("Você escolheu monitorar um servidor Web!")
		fmt.Println("Monitorando servidor Web...")
		// break - Não é necessário, pois o switch case já finaliza após o comando ser executado
	case 2:
		fmt.Println("Você escolheu monitorar um servidor de Banco de Dados!")
		fmt.Println("Monitorando servidor de Banco de Dados...")
	case 3:
		fmt.Println("Obrigado por usar o programa de monitoramento de servidores!")
		fmt.Print("Saindo do programa...")
	default:
		fmt.Println("Comando inválido! Por favor, selecione uma opção válida.")
	}

	// fmt.Scanf("%d", &comando) - Leitura do comando do usuário e salva na variável comando (alocado no endereço de memória)
	//& endereço da variável comando, para que o valor lido seja armazenado nela
}
