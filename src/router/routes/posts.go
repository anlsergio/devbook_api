package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI: "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI: "/posts",
		Method: http.MethodGet,
		Function: controllers.GetPosts,
		RequiresAuth: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodGet,
		Function: controllers.GetPost,
		RequiresAuth: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		RequiresAuth: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		RequiresAuth: true,
	},
	{
		URI: "/posts/{postID}/like",
		Method: http.MethodPost,
		Function: controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI: "/posts/{postID}/dislike",
		Method: http.MethodPost,
		Function: controllers.DislikePost,
		RequiresAuth: true,
	},
}