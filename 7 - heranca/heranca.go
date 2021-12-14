package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Theo", "Ribeiro", 29, 171}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UFPB"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.faculdade)

}
