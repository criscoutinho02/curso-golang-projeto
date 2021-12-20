package main

import "math"

//Inicializando com letra maiúscula é PUBLICO(visivel dentro e fora do pacote)
//Inicializando com letra minuscula é PRIVADO(visivel apenas dentro do pacote)

//Por exemplo
//Ponto - gerará lgo publico
//ponto ou _Ponto gerará algo privado

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distancia responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
