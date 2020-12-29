package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// SetRouter configures all routes inside the router
func SetRouter(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.CheckAuth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}