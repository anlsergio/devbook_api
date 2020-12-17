package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create returns a router containing all the configured routes
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.SetRouter(r)
}