package main

import (
	"fmt"
)

func main() {
	//01. Create a go program that prints the following(up to 100 characters):
	// A
	// AA
	// AAA
	// AAAA
	// AAAAA
	// AAAAAA
	// AAAAAAA
	// ...

	str := "A"

	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", str)
		str += "A"
	}
}
