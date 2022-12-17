package main

import (
	"fmt"
	"time"
)

func main() {
	canal, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		msg1 := <-canal
		msg2 := <-canal2

		fmt.Println(msg1)
		fmt.Println(msg2)
	}
}
