package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{
		modelo:          "f40",
		turboLigado:     false,
		velocidadeAtual: 0,
	}

	var carro2 esportivo = &ferrari{
		modelo:          "f40",
		turboLigado:     false,
		velocidadeAtual: 0,
	}

	fmt.Println(carro1)
	fmt.Println(carro2)

	carro1.ligarTurbo()
	carro2.ligarTurbo()

	fmt.Println(carro1)
	fmt.Println(carro2)
}
