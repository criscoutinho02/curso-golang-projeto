package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo => ", a%b)

	//bitwise
	fmt.Println("E => ", a&b)
	fmt.Println("Ou => ", a|b)
	fmt.Println("Xor => ", a^b)

	c := 2.0
	d := 3.0

	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Mínimo => ", math.Min(c, d))
	fmt.Println("Exponencial => ", math.Pow(c, d))

}
