package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (usr usuario) salvar() {
	fmt.Println("Metodo Salvar")
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", usr.nome)
}

func (usr usuario) maiorDeIdade() bool {
	return usr.idade >= 18
}

func (usr *usuario) fazerAniversario() {
	usr.idade++
}

func main() {
	usr1 := usuario{"Theo", 20}
	fmt.Println(usr1)

	usr1.salvar()

	usr2 := usuario{nome: "Joao", idade: 19}
	usr2.salvar()

	maiorDeIdade := usr2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usr2.fazerAniversario()
	fmt.Println(usr2.idade)
}
