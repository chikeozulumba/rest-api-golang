package router

import (
	"api/router/routes"
	"github.com/gorilla/mux"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	routes.SetupRoutesWithMiddleware(r)
	return routes.SetupRoutes(r)
}