package routes

import "net/http"

// Route represent all routes API
type Route struct {
  URI string
  Method string
  Func func(http.ResponseWriter, *http.Request)
  RequestAuth bool
}


