package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
  var waitGroup sync.WaitGroup

  waitGroup.Add(2) //*! Define quantas goroutines serão executadas

go func() {
    escrever("Olá Mundo")
    waitGroup.Done()
  }()

go func() {
 	  escrever("Programando em Go!")
    waitGroup.Done()
  }()

  waitGroup.Wait()
}

func escrever(texto string) {
  for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
  fmt.Println("!!!! Processo Finalizado !!!!")
}
