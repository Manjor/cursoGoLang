package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo pois os nomes dos retornos ja foram atribuidos
}

func main() {
	seg, pri := trocar(10, 20)
	fmt.Println(seg, pri)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
