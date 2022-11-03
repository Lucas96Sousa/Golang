package main

import (
	"fmt"
	"html"
)

/*
! Fazendo o redirecionamento com os canais "destino recebendo origem"
*/
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

/*
* Multiplexar - misturar (mensagens) em um canal
 */

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.amazon.com.br", "www.uol.com.br"),

		html.Titulo("https://www.globo.com", "www.terra.com.br"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
