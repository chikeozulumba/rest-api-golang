package routes

import (
	"api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		Uri: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
		AuthRequired: true,
	},
	Route{
		Uri: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		AuthRequired: false,
	},
	Route{
		Uri: "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
		AuthRequired: true,
	},
	Route{
		Uri: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		AuthRequired: true,
	},
	Route{
		Uri: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		AuthRequired: true,
	},
}

