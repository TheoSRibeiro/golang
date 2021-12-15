package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops\n")

	i := 0
	/*
		 for i < 10 {
			i++
			fmt.Println("Incrementando i")
			time.Sleep(time.Second)
		}
		fmt.Println(i)

		for j := 0; j < 10; j++ {
			fmt.Println("Incrementando j - ", j)
			time.Sleep(time.Millisecond)
		} */

	nomes := [3]string{"joao", "maria", "jose"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, nome := range "PALAVRA" {
		fmt.Println(indice, nome)
	}

	for indice, nome := range "PALAVRA" {
		fmt.Println(indice, string(nome))
	}

	usr := map[string]string{
		"nome":      "Leo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usr {
		fmt.Println(chave, valor)
	}

	for {
		i++
		fmt.Println("loop infinito.. ", i)
		time.Sleep(time.Second)
	}

}
