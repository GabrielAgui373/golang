package main

import (
	"fmt"
	"math"
)

func main() {
	//A funçaõ de conversão de tipos é do tipo T(v), onde v é convertido para o tipo T
	i := -3
	f := float64(i)
	u := uint(f)

	fmt.Printf("(type: %T value: %v) (type: %T value: %v) (type: %T value: %v)\n", i, i, f, f, u, u)

	//Em Go, uma atribuição de tipos diferentes requer uma conversão explícita, diferente do C
	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f2)

	fmt.Println(f2, x, y, z)
}
