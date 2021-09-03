package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(3200)) // reflect typeOf mostra o tipo

	// sem sinal uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	// com sinal
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa o mapeamento da tabela unicode ( int32 )
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais ( float32 , float64 )
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é ", reflect.TypeOf(49.99)) // Padrão float64

	// booleano
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// strings
	s1 := "Ola meu nome é manoel"
	fmt.Println("O tipo de s1 é ", reflect.TypeOf(s1))
	fmt.Println("O tamanho de s1 é de ", len(s1))
	fmt.Println(s1 + "!")

	// string multline
	s2 := `Ola 
	meu nome 
	é manoel`

	fmt.Println(s2)
	fmt.Println(len(s2))

	char := 'a'
	fmt.Println("O tipo de char é ", reflect.TypeOf(char))
	fmt.Println(char)
}
