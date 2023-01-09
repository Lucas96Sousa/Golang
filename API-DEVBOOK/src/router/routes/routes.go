package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/src/middlewares"
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
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequestAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}

	}

	return r
}
