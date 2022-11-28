package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()

	return c
}

func joint(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := joint(falar("Juan"), falar("Maria"))
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
