package main

import (
	"fmt"
	"runtime"
	"time"
)

func os(os string) string {
	//O switch do Go executa somente o caso selecionado. A declaração break é fornecido automaticamente, diferente de outras linguagens
	switch os {
	case "darwin":
		return "OS X"
	case "Linux":
		return "Linux"
	default:
		return os
	}
}

func say(hour int) string {
	//utilizando condicional nos case
	switch {
	case hour < 12:
		return "Good Morning!"
	case hour < 17:
		return "Good Afternoon!"
	default:
		return "Good Evening!"
	}
}

func main() {
	fmt.Println(os(runtime.GOOS))
	fmt.Println(say(time.Now().Hour()))
}
