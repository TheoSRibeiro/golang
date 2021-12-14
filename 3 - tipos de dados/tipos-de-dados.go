package main

import (
	"errors"
	"fmt"
)

func main() {

	//INT
	var numero int64 = -100000000
	fmt.Println(numero)

	var numero2 uint64 = 10000000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//FLOAT
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123334455555.32
	fmt.Println(numeroReal2)

	//uso de inferencia
	numeroReal3 := 12345.455
	fmt.Println(numeroReal3)

	//STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	//captura o valor da tabela ascii
	char := 'A'
	fmt.Println(char)

	//variavel = 0
	var num int16
	fmt.Println(num)

	var texto string
	fmt.Println(texto)

	//BOOLEAN
	var booleano1 bool
	fmt.Println(booleano1)

	//error
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
