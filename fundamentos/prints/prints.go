package main

import "fmt"

func main() {

	fmt.Print("texto livre")
	fmt.Print("mesma linha")

	fmt.Println("Nova linha")
	fmt.Println("Depois Nova linha")

	x := 3.141516

	//fmt.Println("O valor de x é" + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)
	fmt.Printf("O valor de x é %.2f.", x)

	a, b, c, d := 1, 1.999, false, "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	//("%d %f %.1f %t %s", a, b, c, d)

}
