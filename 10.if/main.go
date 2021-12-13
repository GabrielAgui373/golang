package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//if simples em Go
	x := 10
	if x > 10 {
		fmt.Println(x)
	}

	//if e switch aceitam uma instrução de inicialização muito utilizada para inicializar uma variável local
	if y := x; y < 10 {
		//...
	}

	//Quando uma instrução if não flui para proxima instrução, ou seja, o corpo termina em break, continue ou return, o else desnecessário é omitido
	err, file := open()
	if err == nil {
		fmt.Println(file.Size())
	}
}

func open() (error, os.FileInfo) {
	f, err := os.Open("path")
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	d, err := f.Stat()
	if err != nil {
		f.Close()
		log.Fatal(err)
		return err, nil
	}

	return nil, d
}
