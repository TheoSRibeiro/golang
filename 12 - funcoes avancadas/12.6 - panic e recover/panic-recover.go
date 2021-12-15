package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}

	panic("A média é 7")
}

func main() {
	fmt.Println(alunoAprovado(7, 7))
	fmt.Println("Pós execução!")
}
