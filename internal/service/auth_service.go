package service

import (
	"errors"
	"todo-api/internal/models"
	"todo-api/internal/repository"
	"todo-api/internal/utils"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(r *repository.UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Register(username, email, password string) error {
	// check if exist
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hash,
	}

	return s.repo.Create(&user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
