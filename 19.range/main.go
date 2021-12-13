package main

import "fmt"

func main() {
	//range itera sobre os elementos de diferentes estruturas de dados

	//quando itera sobre um slice, o índice e uma cópia do elemento naquele índice é retornado
	pow := []int{1, 2, 4, 8, 16, 32, 64}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	//utilizando blank identifier
	sum := 0
	for _, num := range pow {
		sum += num
	}
	fmt.Println(sum)
}
