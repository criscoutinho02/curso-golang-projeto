package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[1234567890] = "Maria"
	aprovados[98765432] = "Pedro"
	aprovados[345678] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	for _, nome := range aprovados {
		fmt.Printf("%s \n", nome)
	}

	fmt.Println(aprovados[1234567890])

	delete(aprovados, 1234567890)

	fmt.Println("deletado", aprovados[1234567890])
}
