package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	routes := GetRoutes()
	for _, route := range routes {
		if len(route.Middlewares) > 0 {
			r.Handle(route.Method, route.Path, append(route.Middlewares, route.Handler)...)
		} else {
			r.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
