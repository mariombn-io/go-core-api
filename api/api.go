package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariombn-io/mariombn-io/api/handlers"
	"github.com/mariombn-io/mariombn-io/api/middlewares"
)

type Route struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func GetRoutes() []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/api/auth",
			Handler: handlers.Auth,
		},
		{
			Method:  "GET",
			Path:    "/api/me",
			Handler: handlers.Me,
			Middlewares: []gin.HandlerFunc{
				middlewares.JWTMiddleware,
			},
		},
	}
}
