package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Fn           func(http.ResponseWriter, *http.Request)
	Authenticate bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Fn).Methods(route.Method)
	}

	return r
}
