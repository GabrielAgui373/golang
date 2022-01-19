package main

import (
	"fmt"
	"sync"
)

//Um Channel é utilizado para realizar cominicação bidirecional ou unidirecional entre Goroutines. Somenete dados de um mesmo tipo pode transitar em um Channel

func multiplyByTwo(wg *sync.WaitGroup, ch chan int, num int) {
	defer wg.Done()
	ch <- 2 * num
}
func main() {
	//Assim slice e map, um channel é uma refeência para uma estrutura de dados subjacente
	ch := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go multiplyByTwo(wg, ch, 5)
	fmt.Println(<-ch)

	//Channel com Buffer: um channel com buffer é um channel que possui capacidade de armazenamento dentro dele.
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2

	fmt.Println(len(ch1), cap(ch1))

	wg.Wait()

}
