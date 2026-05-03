package routes

import (
	"todo-api/internal/handler"
	"todo-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	authHandler := handler.NewAuthHandler(db)

	api := r.Group("/api")

	// public
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// protected
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/test", authHandler.Test)

	todoHandler := handler.NewTodoHandler(db)
	protected.POST("/todos", todoHandler.CreateTodo)
	protected.GET("/todos", todoHandler.GetTodos)
	protected.GET("/todos/:id", todoHandler.GetTodoByID)
	protected.PATCH("/todos/:id", todoHandler.UpdateTodo)
	protected.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
