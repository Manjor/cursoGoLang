package main

import "fmt"

// Não tem operadores ternarios
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

// nota >= 6 ? "Aprovado" : "Reprovado" - Uso normal do ternário

func main() {
	fmt.Println(obterResultado(6.2))
}
