package main

import "fmt"

func main() {
	//homogeneo mesmo tipo e estatitico
	var notas [3]float64

	notas[0], notas[1], notas[2] = 7.8, 9.4, 5.7
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

}
