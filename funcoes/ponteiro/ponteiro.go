package main

import "fmt"

func inc1(n int) {
	n++
}

//ponteiro é representado por *
func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n) //por valor
	fmt.Println(n)
	//revisão & obtem o endereço da variavel
	inc2(&n)
	fmt.Println(n)

}
