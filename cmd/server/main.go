package main

import (
	_ "todo-api/docs"
	"todo-api/internal/database"
	"todo-api/internal/routes"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := database.Connect()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}