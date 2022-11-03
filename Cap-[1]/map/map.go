package main

import "fmt"

func main() {
	// var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[49264993630] = "Maria"
	aprovados[10691925208] = "Pedro"
	aprovados[27298464418] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[49264993630])
	delete(aprovados, 49264993630)
	fmt.Println(aprovados[49264993630])
}
