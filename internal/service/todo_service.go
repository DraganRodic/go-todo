package service

import (
	"todo-api/internal/models"
	"todo-api/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(title string, userID uint) (*models.Todo, error) {
	todo := &models.Todo{
		Title:     title,
		Completed: false,
		UserID:    userID,
	}

	err := s.repo.Create(todo)
	return todo, err
}

func (s *TodoService) GetTodos(userID uint) ([]models.Todo, error) {
	return s.repo.GetByUserID(userID)
}

func (s *TodoService) GetTodoByID(userID, id uint) (*models.Todo, error) {
	return s.repo.GetByID(userID, id)
}

func (s *TodoService) UpdateTodo(userID, id uint, title *string, completed *bool) (*models.Todo, error) {
	todo, err := s.repo.GetByID(userID, id)
	if err != nil {
		return nil, err
	}

	if title != nil {
		todo.Title = *title
	}

	if completed != nil {
		todo.Completed = *completed
	}

	err = s.repo.Update(todo)
	return todo, err
}

func (s *TodoService) DeleteTodo(userID, id uint) error {
	todo, err := s.repo.GetByID(userID, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(todo)
}

func (s *TodoService) GetTodosAdvanced(
	userID uint,
	limit, offset int,
	completed *bool,
	sort string,
) ([]models.Todo, int64, error) {
	return s.repo.GetAdvanced(userID, limit, offset, completed, sort)
}
