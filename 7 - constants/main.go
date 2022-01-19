package main

import "fmt"

func main() {
	//Contantes não podem ser declatadas usando a sintaxe ':='
	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")

	const World = "世界"
	fmt.Println("Hello", World)
}
