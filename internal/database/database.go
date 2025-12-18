package database

import (
	"fmt"
	"log"

	"go-rest-api/internal/config"
	"go-rest-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}

func Migrate() error {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}