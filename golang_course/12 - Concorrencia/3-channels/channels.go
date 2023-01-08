package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //*!Definindo qual é o tipo do canal e seus valores de recebimento
	go escrever("Lorem Ipsum", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //*!Define o valor que o canal irá receber -> No caso é o texto
		time.Sleep(time.Second)
	}
	fmt.Println("[---!!!!---Processo Finalizado---!!!!----]")

	close(canal)
}
