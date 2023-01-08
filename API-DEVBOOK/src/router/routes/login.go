package routes

import "net/http"

var routeLogin = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Func:        func(w http.ResponseWriter, r *http.Request) {},
	RequestAuth: false,
}
