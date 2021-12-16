package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!") //goroutine -> concorrencia - nao espera terminar para passar para proxima linha
	escrever("...")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
