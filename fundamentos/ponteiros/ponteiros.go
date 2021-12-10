package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmetica de ponteiros
	//p++

	var p *int = nil
	p = &i //pegando o endereço de i e apontando a p
	*p++   //valor associoado ao ponteiro aumenta 1
	fmt.Println(p, *p, i, &i)

}
