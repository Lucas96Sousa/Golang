package main

import (
	"fmt"

	"github.com/badoux/checkmail"

	"module/auxiliar"
)

func main() {
	fmt.Println("checkmail")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("lucas996oliveira@gmail.com")

	fmt.Println(erro)

}
