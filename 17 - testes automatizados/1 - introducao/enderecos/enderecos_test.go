package enderecos_test

import "testing"
import . "introducao-testes/enderecos" // o '.' faz alias para nao ter que digitar enderecos.TipoEndereco

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TESTE UNITARIO
func TestTipoEndereco(t *testing.T) {
	t.Parallel()

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

}
func TestQualquer(t *testing.T) {
	t.Parallel()

	if 2 < 1 {
		t.Errorf("Teste quebrou!")
	}
}
