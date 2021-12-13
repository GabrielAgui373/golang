package main

import "fmt"

func main() {
	//Slices são abstrações do Array e permite ser dinamicamente dimensionado. Na prática, são mais comuns que arrays
	//Slice é da forma []T, onde T é o tipo dos elementos
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slice literal
	t := []bool{true, false, true} //Aqui temos a criação de um array [3]bool e em seguida uma slice que faz referência a esse array
	fmt.Println(t)

	//a[low(incluso) : high(excluso)]
	var a []int = primes[0:4]
	fmt.Println(a)

	//Slice descreve um array subjacente, logo slices que partilham da mesma referência serão afetadas caso o array subjacente seja alterado
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	b := names[0:2]
	c := names[1:3]

	c[0] = "XXX"
	fmt.Println(b)
	fmt.Println(names)

	//Expressões equivalentes
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s[0:10])
	fmt.Println(s[:10])
	fmt.Println(s[0:])
	fmt.Println(s[:])

	//Comprimento e Capacidade: comprimento é o número de elementos de uma slice e a capacidade é o número de elementos no array subjacente
	d := []int{2, 3, 5, 7, 11, 13}
	printSlice(d)

	d = d[:0]
	printSlice(d)

	d = d[:3]
	printSlice(d)

	d = d[2:]
	printSlice(d)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
