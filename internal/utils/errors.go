package utils

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewBadRequest(msg string) *AppError {
	return &AppError{Code: 400, Message: msg}
}

func NewUnauthorized(msg string) *AppError {
	return &AppError{Code: 401, Message: msg}
}

func NewNotFound(msg string) *AppError {
	return &AppError{Code: 404, Message: msg}
}

func NewInternal(msg string) *AppError {
	return &AppError{Code: 500, Message: msg}
}