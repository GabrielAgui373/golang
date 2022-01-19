package main

import (
	"fmt"
	"time"
)

//A instrução Select é equivalente a instrução Switch, porém o case refere-se ao enviar e recebimento de dados em um Channel. Essa instrução pode ser utilizada quando temos vários Channels esperando para receber informações e queremos executar uma ação quando qualquer um deles for concluído primeiro.

func fast(num int, ch chan int) {
	time.Sleep(5 * time.Millisecond)
	ch <- 2 * num
}

func slow(num int, ch chan int) {
	time.Sleep(15 * time.Millisecond)
	ch <- 2 * num
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go fast(2, ch1)
	go slow(3, ch2)

	select {
	case res := <-ch1:
		fmt.Println("fast finished first, result: ", res)
	case res := <-ch2:
		fmt.Println("slow finished first, result", res)
	}
}
