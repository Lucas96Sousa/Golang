package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all routes API
type Route struct {
	URI         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

// Conf insert all routes in router
func Conf(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
