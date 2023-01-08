package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pet struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	petJSON := `{
      "nome":"Ted",
      "raca":"SRD",
      "idade":6
    }`
	var c pet

	if err := json.Unmarshal([]byte(petJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	pet2JSON := `{"nome":"Bob","raca":"SRD"}`

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(pet2JSON), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}
