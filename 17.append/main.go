package main

import "fmt"

func main() {
	//Para adicionar novos elementos em um slice, Go dispõe da função append
	//func append(s []T, vs ...t) []T - a função recebe como parâmatros o slice e os valore que que serão adicionados

	//obs: se o array subjacente for pequeno, um novo array será alocado e o slice retornado apontará para esse novo array alocado

	var a []int
	printSlice(a)

	a = append(a, 0)
	printSlice(a)

	a = append(a, 1)
	printSlice(a)

	a = append(a, 2, 3, 4)
	printSlice(a)

}

func printSlice(a []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
