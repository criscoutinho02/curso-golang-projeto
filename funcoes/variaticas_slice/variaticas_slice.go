package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {

	aprovados := []string{"Rocky", "Lola", "Panda", "Ariel", "Ayka"}
	imprimirAprovados(aprovados...)

}
