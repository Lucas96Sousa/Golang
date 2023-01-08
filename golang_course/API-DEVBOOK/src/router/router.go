package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Routes Return routes
func Routes() *mux.Router {
	r := mux.NewRouter()
	return routes.Conf(r)
}
