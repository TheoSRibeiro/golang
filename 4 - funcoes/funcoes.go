package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Texto 1")

	var g = func(txt string) string {
		//fmt.Println(txt)
		return txt
	}
	resultado := g("Texto 2")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(10, 5)
	fmt.Println(resultadoSoma2)

}
