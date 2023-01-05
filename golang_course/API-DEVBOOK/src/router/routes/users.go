package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
 {
    URI: "/users",
    Method: http.MethodPost,
    Func: controllers.CreateUser,
    RequestAuth: false,
  },
{
    URI: "/users",
    Method: http.MethodGet,
    Func: controllers.FindUsers,
    RequestAuth: false,
  },
 {
    URI: "/users/{userId}",
    Method: http.MethodGet,
    Func: controllers.FindUserbyId,
    RequestAuth: false,
  }, 
  {
    URI: "/users/{userId}",
    Method: http.MethodPut,
    Func: controllers.UpdateUser,
    RequestAuth: false,
  },
{
    URI: "/users/{userId}",
    Method: http.MethodDelete,
    Func: controllers.DeleteUser,
    RequestAuth: false,
  },
}
