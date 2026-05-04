package repository

import (
	"todo-api/internal/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) GetByUserID(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.DB.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(userID, id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error
	return &todo, err
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.DB.Save(todo).Error
}

func (r *TodoRepository) Delete(todo *models.Todo) error {
	return r.DB.Delete(todo).Error
}

func (r *TodoRepository) GetAdvanced(
	userID uint,
	limit, offset int,
	completed *bool,
	sort string,
) ([]models.Todo, int64, error) {

	var todos []models.Todo
	var total int64

	query := r.DB.Model(&models.Todo{}).Where("user_id = ?", userID)

	// filtering
	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	// count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting
	order := "created_at desc"
	if sort == "asc" {
		order = "created_at asc"
	}

	// fetch
	err := query.
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&todos).Error

	return todos, total, err
}
