package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Arquivo Structs")

	var usr usuario

	usr.nome = "Theo"
	usr.idade = 29
	fmt.Println(usr)

	enderecoEx := endereco{"Rua ali", 1}

	usr2 := usuario{"Theo", 29, enderecoEx}
	fmt.Println(usr2)

	usr3 := usuario{nome: "Theo"}
	fmt.Println(usr3)

	usr4 := usuario{idade: 29}
	fmt.Println(usr4)
}
