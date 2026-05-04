package database

import (
	"log"

	"todo-api/internal/config"
	"todo-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	cfg := config.AppConfig

	dsn := cfg.DBUser + ":" + cfg.DBPass +
		"@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" +
		cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// AUTO MIGRATE
	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	log.Println("Database connected & migrated")

	return db
}
