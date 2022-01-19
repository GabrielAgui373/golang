package main

import "fmt"

func main() {
	const str string = "Hello World"
	fmt.Println(str)

	//Quando se usa aspas simples, o valor armazenado é o valor do caractere na tabela ASCII
	char := 'a'
	fmt.Printf("Type: %T Value: %v", char, char)

	
}
