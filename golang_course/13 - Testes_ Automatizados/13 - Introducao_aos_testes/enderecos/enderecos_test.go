package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do endereco esperado")
	}
}
