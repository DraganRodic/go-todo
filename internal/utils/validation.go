package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fe := range ve {
			out[fe.Field()] = fe.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"validation_error": out,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "invalid request body",
	})
}