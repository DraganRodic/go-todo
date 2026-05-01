package main

import (
	_ "todo-api/docs"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// Ping godoc
// @Summary Ping test
// @Description test endpoint
// @Tags test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", Ping)

	r.Run(":8080")
}