package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("canal_1"), escrever("canal_2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//recebe 2 canais como entrada e retorna 1 como saida
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() { //goroutine anonima
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
