package main

import (
	"fmt"

	"intro-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Rua Candido Rodrigues")

	fmt.Println(tipoEndereco)
}
