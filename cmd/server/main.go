package main

import (
	_ "todo-api/docs"
	"todo-api/internal/config"
	"todo-api/internal/database"
	"todo-api/internal/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadConfig()

	db := database.Connect()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, db)

	r.Run(":" + config.AppConfig.Port)
}
