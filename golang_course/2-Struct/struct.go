package main

import "fmt"

//* faz referencia com os interfaces do typescript
type user struct {
  nome string
  idade uint8
}

func main() {
  fmt.Println("Archive struct")
  
  var u user
  u.nome ="Lucas"
  u.idade=21
  fmt.Println(u)


}
