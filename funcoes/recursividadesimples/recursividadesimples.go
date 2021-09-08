package main

import "fmt"

func fatoria(n uint) uint {

	switch {
	case n == 0:
		return 1
	default:
		return n * fatoria(n-1)
	}

}

func main() {
	resultado := fatoria(5)
	fmt.Println(resultado)
}
