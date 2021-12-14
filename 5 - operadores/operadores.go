package main

import (
	"fmt"
)

func main() {

	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 3 * 4
	restoDivisao := 2 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 2
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUICAO
	var var1 string = "String"
	var var2 string = "String2"
	fmt.Println(var1, var2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	fmt.Println("-- Operadores logicos --")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNARIOS
	fmt.Println("-- Operadores Unarios --")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 20
	fmt.Println(numero)
	numero *= 2
	fmt.Println(numero)
	numero /= 10
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	//OPERADORES TERNARIOS
	// texto := numero > 5 ? "maior que 5" : "menor que 5"

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}
