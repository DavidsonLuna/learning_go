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
	fmt.Scanf("%d", &command)

	fmt.Println("The chosen command was", command)

	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name))
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(age))
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(version))
	fmt.Println("O endereço da minha variavel command", &command)
}
