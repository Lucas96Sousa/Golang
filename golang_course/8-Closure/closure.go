package main

import "fmt"

func closure() func() {
	text := "Inside function closure"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Inside function main"

	fmt.Println(text)

	funcNew := closure()
	funcNew()

}
