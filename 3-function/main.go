package main

import (
	"fmt"
	"math"
)

func main() {
	sum := sum(10, 2)
	fmt.Println(sum)

	fmt.Println("-----High Order Function------")

	//Passando uma função como parãmetro
	Sphare(func(radius float64) float64 {
		return (4 / 3) * (math.Pi * radius * radius * radius)
	})

	//Retornando uma função
	double := getMultiplier(2)
	triple := getMultiplier(3)
	quadriple := getMultiplier(4)

	fmt.Println(double(3), triple(3), quadriple(3))

	//Multiplos retornos
	sum, sub, _, mult := operation(2, 3)
	fmt.Println(sum, sub, mult)

	//Função Anônima
	square := func(l int) int {
		return l*l
	}
	fmt.Println(square(2))
}

//Função básica
func sum(x, y int) int {
	return x + y
}

func Sphare(vol func(radius float64) float64) {
	fmt.Println("Volume of sphare is:", vol(5))
}

func getMultiplier(multiplier int) func(number int) int {
	return func(number int) int {
		return multiplier * number
	}
}

func operation(num1, num2 int) (sum, sub, div, mult int) {
	sum = num1 + num2
	sub = num1 - num2
	div = num1 / num2
	mult = num1 * num2

	return
}
