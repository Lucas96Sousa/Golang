package main

import "fmt"

// * faz referencia com os interfaces do typescript
type user struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Archive struct")

	var u user
	u.nome = "Lucas"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua xpto", 1}

	user2 := user{"Davi", 21, enderecoExemplo}
	fmt.Println(user2)

}
