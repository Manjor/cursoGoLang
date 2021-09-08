package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{

		"G": {
			"Gabriel": 1545.2,
			"Guga":    1200.2,
		},
		"J": {
			"Jose": 2340.00,
		},
		"P": {
			"Pedro": 1200.1,
		},
	}

	delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
