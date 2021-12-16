package main

import "fmt"

func fatorial(n uint) uint {
	if n <= 1 {
		return 1
	} else {
		fatorialAnterior := fatorial(n - 1)
		return fatorialAnterior * n
	}

}

func main() {
	resultado := fatorial(10)
	fmt.Println("Resultado: ", resultado)

}
