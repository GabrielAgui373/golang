package main

import "fmt"

//Um channel é uma técnica que permite realizar uma comunicação birecional entre Goroutines. Um channel só pode transferrir dados de um mesmo tipo
func myfunc(ch chan int) {

	fmt.Println(234 + <- ch)
}

func main() {
	var channel chan int

	fmt.Println("Value of the channel: ", channel)
	fmt.Printf("Type of the channel: %T", channel)

	channel2 := make(chan int)

	fmt.Println("\nValue of the channel: ", channel2)
	fmt.Printf("Type of the channel: %T", channel2)

	fmt.Println("\n\nstart Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}
