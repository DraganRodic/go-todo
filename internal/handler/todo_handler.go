package handler

import (
	"net/http"
	"strconv"
	"todo-api/internal/repository"
	"todo-api/internal/service"
	"todo-api/internal/utils"

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
	Title string `json:"title" binding:"required,min=3"`
}

type UpdateTodoRequest struct {
	Title     *string `json:"title" binding:"omitempty,min=3"`
	Completed *bool   `json:"completed"`
}

// @Summary Create todo
// @Tags todos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body CreateTodoRequest true "Todo data"
// @Success 201 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	userID, _ := c.Get("user_id")

	todo, err := h.service.CreateTodo(req.Title, userID.(uint))
	if err != nil {
		utils.Error(c, utils.NewInternal(err.Error()))
		return
	}

	utils.Success(c, http.StatusCreated, todo)
}

// @Summary Get user todos
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param completed query bool false "Filter by completed (true/false)"
// @Param sort query string false "Sort by created_at (asc/desc)"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/todos [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// filtering
	completedStr := c.Query("completed")
	var completed *bool
	if completedStr != "" {
		val, err := strconv.ParseBool(completedStr)
		if err != nil {
			utils.Error(c, utils.NewBadRequest("completed must be true or false"))
			return
		}
		completed = &val
	}

	// sorting
	sort := c.DefaultQuery("sort", "desc") // asc | desc
	if sort != "asc" && sort != "desc" {
		utils.Error(c, utils.NewBadRequest("sort must be 'asc' or 'desc'"))
		return
	}

	todos, total, err := h.service.GetTodosAdvanced(userID.(uint), limit, offset, completed, sort)
	if err != nil {
		utils.Error(c, utils.NewInternal(err.Error()))
		return
	}

	utils.Success(c, http.StatusOK, gin.H{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

// @Summary Get todo by id
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, utils.NewBadRequest("invalid id"))
		return
	}

	todo, err := h.service.GetTodoByID(userID.(uint), uint(id))
	if err != nil {
		utils.Error(c, utils.NewNotFound("todo not found"))
		return
	}

	utils.Success(c, http.StatusOK, todo)
}

// @Summary Update todo (partial)
// @Tags todos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Param data body UpdateTodoRequest true "Todo data"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/todos/{id} [patch]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, utils.NewBadRequest("invalid id"))
		return
	}

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	todo, err := h.service.UpdateTodo(userID.(uint), uint(id), req.Title, req.Completed)
	if err != nil {
		utils.Error(c, utils.NewNotFound("todo not found"))
		return
	}

	utils.Success(c, http.StatusOK, todo)
}

// @Summary Delete todo
// @Tags todos
// @Produce json
// @Security BearerAuth
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, utils.NewBadRequest("invalid id"))
		return
	}

	todo, err := h.service.GetTodoByID(userID.(uint), uint(id))
	if err != nil {
		utils.Error(c, utils.NewNotFound("todo not found or not yours"))
		return
	}

	if err := h.service.DeleteTodo(userID.(uint), todo.ID); err != nil {
		utils.Error(c, utils.NewInternal("failed to delete"))
		return
	}

	utils.Success(c, http.StatusOK, gin.H{
		"message": "todo deleted",
	})
}
