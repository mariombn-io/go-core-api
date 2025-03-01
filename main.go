package main

import (
	"fmt"
	"os"

	// Importe seus pacotes reais conforme a estrutura do seu go.mod.
	// Exemplo:
	"github.com/mariombn-io/go-core-api/api"
	"github.com/mariombn-io/go-core-api/cmd/cli"
	// "github.com/mariombn-io/go-core-api/config"
)

func main() {
	// Se quiser inicializar configurações globais (ler .env, etc.), faça aqui:
	// config.Init()

	// Valores padrão
	host := "0.0.0.0"
	port := "8088"

	// Se não houver parâmetros, mostra ajuda
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Flag para saber se o modo é API ou CLI
	mode := ""

	// Processa argumentos manualmente
	// Exemplo de uso:
	//   go run main.go --api --host 127.0.0.1 --port 3000
	//   go run main.go -a -h 127.0.0.1 -p 3000
	//   go run main.go --cli
	i := 1
	for i < len(os.Args) {
		arg := os.Args[i]
		switch arg {
		case "--help":
			printHelp()
			return
		case "--api", "-a":
			mode = "api"
		case "--cli", "-c":
			mode = "cli"
		case "--host", "-h":
			if i+1 < len(os.Args) {
				i++
				host = os.Args[i]
			} else {
				fmt.Println("Erro: faltou valor para a flag --host ou -h")
				return
			}
		case "--port", "-p":
			if i+1 < len(os.Args) {
				i++
				port = os.Args[i]
			} else {
				fmt.Println("Erro: faltou valor para a flag --port ou -p")
				return
			}
		default:
			// Se aparecer algo que não reconhecemos, mostramos ajuda.
			fmt.Printf("Flag ou argumento desconhecido: %s\n", arg)
			printHelp()
			return
		}
		i++
	}

	// Se não definiu modo, avisa e sai
	if mode == "" {
		fmt.Println("Nenhuma opção fornecida. Use --api ou --cli.")
		printHelp()
		return
	}

	// Decide o que rodar
	switch mode {
	case "api":
		// Inicia a API com o host e porta selecionados
		api.StartServer(host, port)
	case "cli":
		// Inicia o CLI (apenas imprime algo, por enquanto)
		cli.RunCLI()
	}
}

func printHelp() {
	fmt.Println(`
Usage: go run main.go [opções]

Sem parâmetros: mostra esta ajuda.

Opções:
  --api, -a            Inicia a API
  --cli, -c            Inicia o CLI (Em construção)
  --host, -h <host>    Define o host da API (padrão: 0.0.0.0)
  --port, -p <porta>   Define a porta da API (padrão: 8088)
  --help               Exibe esta mensagem de ajuda

Exemplos:
  go run main.go --api --host 127.0.0.1 --port 3000
  go run main.go -a -h 127.0.0.1 -p 3000
  go run main.go --cli
`)
}
