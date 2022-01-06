//Um programa GO é composto de pacotes e começa rodando pelo pacote main
//Um pacote é formado por todos os arquivos de um diretório
package main

import (
	"fmt"
	"math/rand"
	"module/say"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	//Todos os nomes exportados começam com letra maiúscula. Os nomes "não exportados" não são acessíveis fora do pacote
	say.SayHi("Gabriel")
}
