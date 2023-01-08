package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Func:        controllers.Login,
	RequestAuth: false,
}
