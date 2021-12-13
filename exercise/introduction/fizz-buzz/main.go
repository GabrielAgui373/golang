package main

import "fmt"

func main() {
	//1. Solve this problem, called Fizz-Buzz: write a program that prints the numbers 1 to 100. But for multiples of three print "Fizz" instead to the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print FizzBuzz
	const (
		FIZZ = 3
		BUZZ = 5
	)

	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
