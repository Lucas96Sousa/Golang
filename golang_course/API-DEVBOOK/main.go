package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting API ---> http://localhost:5000")

	r := router.Routes()
	log.Fatal(http.ListenAndServe(":5000", r))
}
