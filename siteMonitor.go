package main

// Multiplos imports
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
const sleep = 5

// Toda função de pacote externo, inicia com letra maiuscula
func main() {
	showIntroduction()

	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			leaveTheProgram()
		default:
			fmt.Println("Comando inválido!")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "Rodolfo"
	version := 1.1

	fmt.Println("Hello sr.", name)
	fmt.Println("This program is in version", version)

}
func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var comando int
	// o caractere & indica que estamos apontando para um endereço de memória
	// o caractere "%d" indica que a entrada do usuário deve ser do tipo inteiro
	// fmt.Scanf("%d", &comando)

	fmt.Scan(&comando)

	return comando
}

func leaveTheProgram() {
	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://www.alura.com.br", "https://www.magix.digital", "https://www.caelum.com.br"}

	sites := readFileSites()

	// o operador _ é conhecido como blank identifier
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			siteTest(site)
		}

		time.Sleep(sleep * time.Second)
	}

}

func siteTest(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println("Ocorreu um erro inesperado:", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
		logRegister(site, false)
	}
}

func readFileSites() []string {
	var sites []string

	file, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Ocorreu um erro inesperado:", error)
	}

	leitor := bufio.NewReader(file)

	for {
		line, err := leitor.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro inesperado:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online: " + strconv.FormatBool(status) + "\n")

}

func showLogs() {
	fmt.Println("Exibindo Logs...")
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
