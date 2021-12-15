package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas")

	somaTotal := soma(10, 21, 23, 45, 76, 933, 423, 47, 63)
	fmt.Println(somaTotal)

	escrever("Olá mundo!", 1, 2, 3, 4, 5, 6, 3)
}
