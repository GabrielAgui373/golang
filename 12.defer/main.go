package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//A instrução defer programa uma chamada para ser executada imediatamente antes do retorno de uma função ou depois da última instrução
	defer fmt.Println("World")
	fmt.Println("Hello")

	slice, content, _ := contents("C:/text.txt")
	fmt.Println(content, slice)
}

//A instrução defer é eficaz para lidar com recursos que devem ser liberados
func contents(filename string) (string, []byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}

	defer f.Close() //O close será chamado quando tudo tiver sido executado

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read((buf[0:]))
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", nil, err
		}

	}

	return string(result), result, nil //Close() será chamado antes do return
}
