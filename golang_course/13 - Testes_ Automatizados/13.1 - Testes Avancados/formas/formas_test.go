package formas

import (
	"math"
	"testing"
)

func TesteArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)

		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %f é diferente da esperada %f", areaEsperada, areaRecebida)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %f é diferente da esperada %f", areaEsperada, areaRecebida)
		}
	})
}
