package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Instancia de rotas do Gin
	r := gin.Default()

	//Rota para o endpoint /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rota para o endpoint /ping
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Instrução",
		})
	})

	// Porta que o servidor irá escutar
	r.Run(":8080")
}
