package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hourCorrect(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") // estilo de formatação do time
	fmt.Fprintf(w, "<h1> Correct hour -> %s </h1>", s)
}

func main() {
	http.HandleFunc("/", hourCorrect)
	log.Println("Starting -> http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
