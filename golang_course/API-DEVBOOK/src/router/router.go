package router

import "github.com/gorilla/mux"

//Routes Return routes 
func Routes() *mux.Router{
  return mux.NewRouter()
}
