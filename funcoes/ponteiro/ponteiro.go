package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * é usado para desreferenciar
	// ter acesso ao valor o qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n)

	inc2(&n) // por referencia
	fmt.Println(n)
}
