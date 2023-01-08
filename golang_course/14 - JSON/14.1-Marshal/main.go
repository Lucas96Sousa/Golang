package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Pet struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := Pet{"Ted", "SRD", 6}

	petJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(petJSON))

	c2 := map[string]string{
		"nome": "Bob",
		"raca": "SRD",
	}
	pet2JSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(pet2JSON))

}
