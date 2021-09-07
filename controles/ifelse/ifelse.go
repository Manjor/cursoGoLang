package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovador com nota: ", nota)
	} else {
		fmt.Println("Reprovado com nota: ", nota)
	}
}

func main() {
	imprimirResultado(4.2)
	imprimirResultado(7.2)
}
