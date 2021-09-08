package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	resultado := soma(10, 12)
	fmt.Println(resultado)

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))
}
