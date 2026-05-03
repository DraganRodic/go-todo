package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}