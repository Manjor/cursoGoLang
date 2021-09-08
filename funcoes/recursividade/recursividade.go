package main

import "fmt"

func fatoria(n int) (int, error) {

	switch {
	case n < 0:
		return -1, fmt.Errorf("numero inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatoria(n - 1)
		return n * fatorialAnterior, nil
	}

}

func main() {
	resultado, _ := fatoria(5)
	fmt.Println(resultado)

	_, err := fatoria(-4)
	if err != nil {
		fmt.Println(err)
	}
}
