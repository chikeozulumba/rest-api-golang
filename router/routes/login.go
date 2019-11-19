package routes

import (
	"api/controllers/auth"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		Uri: "/auth/login",
		Method: http.MethodPost,
		Handler: auth.HandleLogin,
		AuthRequired: false,
	},
}
