package main

import "fmt"

func main() {
	canal := make(chan string, 2) //buffer = 2
	canal <- "canal_1"
	canal <- "canal_2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
