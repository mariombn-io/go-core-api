package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(host, port string) {
	r := gin.Default()

	// Rota de exemplo
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Iniciando servidor em %s\n", address)
	err := r.Run(address)
	if err != nil {
		panic(err)
	}
}
