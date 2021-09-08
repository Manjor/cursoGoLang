package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"Jose Joao": 11325.45,
		"Gabriel":   1000.10,
		"Pedro":     1200.10,
	}

	funcionarios["Rafael"] = 1350.00
	delete(funcionarios, "Inexistente")

	for nome, salario := range funcionarios {
		fmt.Println(nome, salario)
	}
}
