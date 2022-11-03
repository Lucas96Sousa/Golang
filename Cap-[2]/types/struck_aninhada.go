package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 2, preco: 25.40},
			item{produtoID: 2, quantidade: 11, preco: 86.43},
			item{produtoID: 3, quantidade: 300, preco: 3.13},
		},
	}
	fmt.Printf("Valor do pedido é: R$ %.2f", pedido.valorTotal())
}
