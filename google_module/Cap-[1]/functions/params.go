package main

import "fmt"

func main() {
	resultado := exec(multiplicacao, 2, 4)

	fmt.Println(resultado)
}

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
