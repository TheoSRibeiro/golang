package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TESTE UNITARIO
func TestTipoEndereco(t *testing.T) {

	cenariosTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes ABC", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"AVENIDA ABC", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", retornoRecebido, cenario.retornoEsperado)
		}
	}

	/* enderecoTeste := "Rua Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoEndereco(enderecoTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	} */
}
