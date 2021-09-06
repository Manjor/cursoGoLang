package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // type float64
	y := 2

	fmt.Println(x / float64(y)) // convertendo inteiro para floatß

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidados
	fmt.Println("Teste " + string(123)) // string() mostra o valor da abela asc

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // qualquer outro valor alem de 1 ou true não imprime verdadeiro
	if b {
		fmt.Println("Verdadeiro")
	}
}
