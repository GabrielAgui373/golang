package main

import "fmt"

func main() {
	//Quando uma variável é declarada sem um valor inicial, elas recebem um "valor zero" padrão
	var i int //0
	var f float64 //0
	var b bool //false
	var s string // ""

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}