package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("FIM!") // executa antes do retorno da funçao, porem não no inicio
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	resultado := obterValorAprovado(2000)
	fmt.Println(resultado)
}
