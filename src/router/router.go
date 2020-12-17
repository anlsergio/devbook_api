package router

import "github.com/gorilla/mux"

// Create returns a router containing all the configured routes
func Create() *mux.Router {
	return mux.NewRouter()
}