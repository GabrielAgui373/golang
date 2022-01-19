package main

import "fmt"

func main() {
	var (
		a bool = false
		b bool = 3 > 1
	)

	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
}
