package handler

import (
	"net/http"
	"todo-api/internal/repository"
	"todo-api/internal/service"
	"todo-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	repo := repository.NewUserRepository(db)
	service := service.NewAuthService(repo)

	return &AuthHandler{service: service}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// @Summary Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "User data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.service.Register(req.Username, req.Email, req.Password); err != nil {
		utils.Error(c, utils.NewBadRequest(err.Error()))
		return
	}

	utils.Success(c, http.StatusCreated, gin.H{
		"message": "user created",
	})
}

// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body LoginRequest true "Login data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		utils.Error(c, utils.NewUnauthorized("invalid credentials"))
		return
	}

	utils.Success(c, http.StatusOK, gin.H{
		"token": token,
	})
}