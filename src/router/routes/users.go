package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequiresAuth: true,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetUser,
		RequiresAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequiresAuth: false,
	},
}
