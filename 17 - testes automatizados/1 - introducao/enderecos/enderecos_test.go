package enderecos

import "testing"

//TESTE UNITARIO
func TestTipoEndereco(t *testing.T) {
	enderecoTeste := "Rua Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoEndereco(enderecoTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}
