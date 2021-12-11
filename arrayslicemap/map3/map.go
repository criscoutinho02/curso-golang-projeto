package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 2345.99,
			"Guga Pereira":   9876.98,
		},
		"B": {
			"Bruna Silva":   12345.99,
			"Beatriz Silva": 92345.99,
		},
		"P": {
			"Pedro Silva": 567.99,
		},
	}
	delete(funcsPorLetra, "P")

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
