package main

import "fmt"

//A instrução 'var' declara uma lista de variáveis com o último argumeto sendo o tipo
var c, pyhton, java bool

//se o inicializador estiver presente o tipo é inferido de acordo com o inicializador
var pi = 3.1415

func main() {
	var i int
	fmt.Println(c, pyhton, java, i, pi)

	//Declaração curta: a declaração ':=' pode ser utilizada ao invés de 'var'
	say := "Hello World"
	fmt.Println(say)
}
