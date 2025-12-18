package handlers

import (
	"go-rest-api/internal/database"
	"go-rest-api/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// Check handles basic health check
func (h *HealthHandler) Basic(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse("Server is running", fiber.Map{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "go-rest-api",
		"version":   "1.0.0",
	}))
}

// Ready handles readiness check (with database check)
func (h *HealthHandler) Detailed(c *fiber.Ctx) error {
	// Check database connection
	sqlDB, err := database.GetDB().DB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(utils.ErrorResponse("Database connection error"))
	}

	// Ping database
	if err := sqlDB.Ping(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(utils.ErrorResponse("Database ping failed"))
	}

	// Get database stats
	stats := sqlDB.Stats()

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse("Service is ready", fiber.Map{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "go-rest-api",
		"version":   "1.0.0",
		"database": fiber.Map{
			"status":         "up",
			"max_open_conns": stats.MaxOpenConnections,
			"open_conns":     stats.OpenConnections,
			"in_use":         stats.InUse,
			"idle":           stats.Idle,
		},
	}))
}

func (h *HealthHandler) Readiness(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ready": true,
	})
}
