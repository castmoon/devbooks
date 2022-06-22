package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Fn:           controllers.CreateUser,
		Authenticate: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Fn:           controllers.ListUsers,
		Authenticate: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Fn:           controllers.GetUser,
		Authenticate: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Fn:           controllers.UpdateUser,
		Authenticate: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Fn:           controllers.DeleteUser,
		Authenticate: false,
	},
}
