package main

import "fmt"
import "reflect"

func main() {
	var name = "Davidson"
	var age = 40
	var version = 1.1
	fmt.Println("Olá, sr.", name, "sua é idade", age)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name))
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(age))
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(version))
}
