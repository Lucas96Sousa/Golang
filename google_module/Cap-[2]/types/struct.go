package main

import "fmt"

type product struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver(receptor)

func (p product) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var product1 product
	product1 = product{
		nome:     "Lapis",
		preco:    1.83,
		desconto: 0.05,
	}
	fmt.Println(product1, product1.precoComDesconto())

	product2 := product{"Notebook", 1989.90, 0.10}
	fmt.Println(product2.precoComDesconto())
}
