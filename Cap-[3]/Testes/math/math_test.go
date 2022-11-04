package math

import "testing"

const err = "Valor esperado %v, mais o resultado esperado %v"

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperado {
		t.Errorf(err, valorEsperado, valor)
	}
}
