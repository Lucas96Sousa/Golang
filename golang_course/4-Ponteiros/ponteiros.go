package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	//*! Ponteiro é uma referencia de memoria

	var v3 int
	var point *int

	v3 = 100
	point = &v3
	fmt.Println(v3, *point) /*
	  ! desreferenciação de memoria com o *sobre a variavel, irá ser lido o valor interno
	*/
}
