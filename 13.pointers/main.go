package main

import (
	"fmt"
)

func main() {
	//Ponteiros ou apontador armazanam endereço de memória
	//(*) - Operador de diferência
	//(&) - Operador de endereço

	a := 255

	var p *int = &a
	fmt.Println(p)

	//Processo de desreferenciamento: altera o valor que está armazenado no endereço de memória do ponteiro
	*p = 50
	fmt.Println(a)

	//O valor zero de um ponteiro é nil
	var x *int
	fmt.Println(x)
}
   