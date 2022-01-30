package main

import (
	"fmt"
	"sync"
)

//A bilbioteca sync.WaitGroup fornece um contador que bloqueia a execução de uma função até que seu contador interno se torne 0

// .Add(n) -> adiciona n ao contador, indicando o número de Goroutines em execução
// .Done() -> remove 1 do contador, indicando a fim da execução de uma Goroutine
// .Wait() -> Bloqueia a execução até que o contador se torne 0

func run1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am first runner")
}

func run2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am second runner")
}

func exec() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go run1(wg)
	go run2(wg)

	wg.Wait()
}

//Usar waitGroup quando não estiver trabalhando com channels

func main() {
	exec()
}
