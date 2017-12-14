package main

import "fmt"

func main() {
	name := "Forrest"
	version := 1.1
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)

	//option 1 conditional
	// if command == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if command == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if command == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Desculpe, mas não conheço esse comando")

	//option 2 conditional
	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Desculpe, mas não conheço esse comando")
	}
	//Show variables type results
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name))
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(age))
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(version))
	//fmt.Println("O endereço da minha variavel command", &command)
}
