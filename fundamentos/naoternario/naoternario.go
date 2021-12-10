package main

import "fmt"

//nao ha operador ternario

func obtemResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obtemResultado(6.2))
}
