package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(1, 2))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(4, 2))
}
