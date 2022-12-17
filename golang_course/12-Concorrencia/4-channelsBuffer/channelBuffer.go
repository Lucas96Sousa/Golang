package main

import "fmt"

func main(){
  canal := make(chan string, 2 ) //*! Definindo um canal com buffer
  canal <- "Hallo"

  mensagem := <- canal
  fmt.Println(mensagem)

}
