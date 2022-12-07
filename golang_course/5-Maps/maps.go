package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Sousa",
	}

	fmt.Println(user)
}
