package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("canal_1", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("FIM PROGRAMA")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
