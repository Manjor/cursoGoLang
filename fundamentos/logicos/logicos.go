package main

import "fmt"

func comprar(t1, t2 bool) (bool, bool, bool) {
	comprarTv50 := t1 && t2
	comprarTv32 := t1 != t2
	comprarSorvete := t1 || t2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
