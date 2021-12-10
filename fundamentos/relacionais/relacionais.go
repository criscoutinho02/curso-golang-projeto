package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Banana" == "Banana")
	fmt.Println("Strings", 3 != 2)
	fmt.Println("Menor", 3 < 2)
	fmt.Println("Maior", 3 > 2)
	fmt.Println("Menor Igual", 3 <= 2)
	fmt.Println("Maior Igual", 3 >= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Pessoas:", p1 == p2)

}
