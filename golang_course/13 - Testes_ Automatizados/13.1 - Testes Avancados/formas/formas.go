package formas

import (
	"math"
)

// Forma
type Forma interface {
	Area() float64
}

// Retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Circulo
type Circulo struct {
	raio float64
}

// func
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// func
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
