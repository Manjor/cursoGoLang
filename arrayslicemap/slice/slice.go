package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array ! É um trecho de array
	s2 := a2[1:3] // do indice 1 ao 3 sem incluir o 3
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(s3)

	// imaginar o slice como : tamanho e um ponteiro para um elemento de um array
}
