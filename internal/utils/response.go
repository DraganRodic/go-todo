package utils

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, err *AppError) {
	c.JSON(err.Code, gin.H{
		"error": err.Message,
	})
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}