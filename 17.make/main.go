package main

import "fmt"

func main() {
	//A função make aloca um espaço na memória. Nos exemplos abaixo, a função make aloca um array vazio e retorna um slice que faz referência a esse array

	//make(type, len, cap(optional))
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
	
	c := b[:2]
	printSlice(c)
	
	d := c[2:5]
	printSlice(d)


}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
