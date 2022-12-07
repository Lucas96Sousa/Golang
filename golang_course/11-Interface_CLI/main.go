package main

import (
	"fmt"
	"log"
	"os"

	"interface-cli/app"
)

func main() {
	fmt.Println("Starting point...")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
