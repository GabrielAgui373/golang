package main

import (
	"fmt"
	"time"
)

//Goroutine é uma função ou método que é "paralelizada" com outras Goroutine. A ideia é pensar uma Goroutine como uma Thread, que roda em um mesmo espaço de endereço que outras threads em um processo, porém com um custo menor.
func displaty(str string) {
	for w := 0; w < 6; w++ {
		fmt.Printf("\n%s - %d", str, w)
		time.Sleep(time.Second);
	}
}

func main() {
	go displaty("Goroutine")

	displaty("Welcome")

	//Goroutine anônima
	go func() {
		fmt.Println("Anonymous")
	}()

	time.Sleep(time.Second)
}