package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Println(a1)
	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	//slice não é um array , slice é um pedaço de um array
	s2 := a2[1:3]

	fmt.Println(a2)
	fmt.Println(s2)

	s3 := a2[:2]
	fmt.Println(s3)

	//slice é um ponteiro com tamanho e ponteiro para um elemento e estrutura continua
	s4 := s2[:1]
	fmt.Println(s4)

	s5 := a2[3:4]
	fmt.Println(s5)

}
