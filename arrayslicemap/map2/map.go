package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"José João":      11234.99,
		"Gabriela Silva": 1200.99,
		"Bruno Santos":   20000.87,
	}

	funcsESalarios["Cristina Coutinho"] = 25000
	delete(funcsESalarios, "Inexistente")

	fmt.Println(funcsESalarios)
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
