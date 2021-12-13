package main

import "fmt"

func main() {
	//1. Create a simple loop with the for construct. Make it loop 10 times and print out the loop counter with the fmt package
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//2. Rewrite the loop again so that it fills an slice and then prints that array to the screen

	//É interessante utilizar o make para alocar o array subjacente com a capacidade determinada, para que ele não precise alocar um novo array em cada iteração
	mat := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		mat = append(mat, i)
	}
	fmt.Println(mat)
}
