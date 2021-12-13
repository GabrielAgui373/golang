package main

import "fmt"

func main() {
	const e float32 = 2.718281828459045235
	const pi float64 = 3.1415926535

	//Quando é utilizada inferência de tipo, o tipo float armazena o número de bits de acordo com a arquitetura do computador
	num1 := 3.43
	fmt.Printf("Type: %T Value: %v\n", num1, num1)

	//complex64 possui uma parte real float32 e uma parte imaginária float64
	const com1 complex64 = complex(43.22, 4.1112)
	fmt.Printf("Type: %T Value: %v\n", com1, com1)

	//Complex128 possui uma parte real float64 e uma parte imaginária float64
	const com2 complex128 = complex(33243.2, 4.1222)
	fmt.Printf("Type: %T Value: %v\n", com2, com2)

}
