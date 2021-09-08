package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345112978] = "Joao"
	aprovados[12345612978] = "Ze"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])
}
