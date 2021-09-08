package main

import "fmt"

func f1() { // sem parametros e retornos
	fmt.Println("F1")
}

func f2(p1 string, p2 string) { // com parametros
	fmt.Printf("F2: %s %s", p1, p2)
}

func f3() string {
	return "F5" // com retorno
}

func f4(p1, p2 string) string { // com parametro e com retorno
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param 1", "Param 2")

	r3, r4 := f3(), f4("Param 1", "Param 2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5: ", r51, r52)

}
