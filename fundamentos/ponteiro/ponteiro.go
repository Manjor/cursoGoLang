package main

import "fmt"

func main() {

	i := 1

	// Go não tem aritmetica de ponteiros
	// * ponteiro de referencia de memória
	// & endereço de memoria

	var p *int = nil

	p = &i // pegando o endereco da variavel i e atribuindo ao ponteiro
	*p++   // pega o valor associado ao ponteiro, incrementando assim o 1 para 2
	i++    // Também incrementa o valor de i

	fmt.Println(p, *p, i, &i)

}
