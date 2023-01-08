package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
  config.Load()
  fmt.Printf("Starting server [=======> http://localhost:%d <============]", config.Port)

	r := router.Routes()
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
