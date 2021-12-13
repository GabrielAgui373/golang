package main

import "fmt"

func main() {
	//Em Go, um array da forma [n]T, onde é o tamanho e T o tipo
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//Arrays não podem ser redimensionaods
	//Array com tamanho determinado pela quantidade de itens
	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
