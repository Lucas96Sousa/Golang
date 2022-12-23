package enderecos_test

import (	
	. "intro-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
    t.Parallel()

	cenariosTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Alameda Santos", "Alameda"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça qtas", "Tipo Inválido"},
		{"RUA NORTE", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosTeste {
		retornoEsperado := TipoEndereco(cenario.enderecoInserido)
		if retornoEsperado != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoEsperado,
				cenario.retornoEsperado,
			)
		}
	}
}
func TestQualquer(t *testing.T) {
  if 13 > 1 {
    t.Errorf("Teste Quebrou")
  }
}
