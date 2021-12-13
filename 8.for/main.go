package main

import "fmt"

func main() {
	//Go possui somente uma estrutura de repetiçãp. A estrutura 'for' e 'white' são unificadas e não há 'do while'
	//Como for no C: for init; codition; post {}
	//Como while no C: for condition {}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//'while' é escrito como for em Go
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//Utilizando a cláusula range
	for _, value := range "World" {
		fmt.Println(value, string(value))
	}
}
