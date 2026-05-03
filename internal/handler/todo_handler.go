package handler

import (
	"net/http"
	"strconv"
	"todo-api/internal/repository"
	"todo-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	repo := repository.NewTodoRepository(db)
	service := service.NewTodoService(repo)

	return &TodoHandler{service: service}
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

// CreateTodo godoc
// @Summary Create todo
// @Tags todos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body CreateTodoRequest true "Todo data"
// @Success 201 {object} map[string]interface{}
// @Router /api/todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	todo, err := h.service.CreateTodo(req.Title, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodos godoc
// @Summary Get user todos
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Todo
// @Router /api/todos [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")

	todos, err := h.service.GetTodos(userID.(uint))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

// GetTodoByID godoc
// @Summary Get todo by id
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /api/todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.service.GetTodoByID(userID.(uint), uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(200, todo)
}

// UpdateTodo godoc
// @Summary Update todo (partial)
// @Tags todos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Param data body UpdateTodoRequest true "Todo data"
// @Success 200 {object} models.Todo
// @Router /api/todos/{id} [patch]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.UpdateTodo(userID.(uint), uint(id), req.Title, req.Completed)
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(200, todo)
}

// DeleteTodo godoc
// @Summary Delete todo
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]string
// @Router /api/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.service.GetTodoByID(userID.(uint), uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "todo not found or not yours"})
		return
	}

	err = h.service.DeleteTodo(userID.(uint), todo.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(200, gin.H{
		"message": "todo deleted",
	})
}
