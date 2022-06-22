package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//return a router with all configured routes
func BuildRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
