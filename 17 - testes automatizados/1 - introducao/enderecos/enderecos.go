package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}
