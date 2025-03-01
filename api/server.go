package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(host, port string) {
	r := gin.Default()
	SetupRoutes(r)
	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Starting server at %s\n", address)
	if err := r.Run(address); err != nil {
		panic(err)
	}
}
