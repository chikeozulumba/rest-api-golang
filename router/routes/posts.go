package routes

import (
	"api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	Route{
		Uri: "/posts",
		Method: http.MethodPost,
		Handler: controllers.CreatePost,
		AuthRequired: true,
	},
	Route{
		Uri: "/posts",
		Method: http.MethodGet,
		Handler: controllers.GetAllPosts,
		AuthRequired: true,
	},
	Route{
		Uri: "/posts/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetPostByID,
		AuthRequired: true,
	},
	Route{
		Uri: "/posts/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdatePost,
		AuthRequired: true,
	},
	Route{
		Uri: "/posts/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeletePost,
		AuthRequired: true,
	},
}

