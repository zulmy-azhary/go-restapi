package routes

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, h *handlers.Handlers, jwtSecret string) {
	api := app.Group("/api")

	// Health Check Routes
	health := api.Group("/health")
	health.Get("/", h.Health.Basic)
	health.Get("/detailed", h.Health.Detailed)
	health.Get("/ready", h.Health.Readiness)

	// Auth Routes (Public)
	auth := api.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.Login)

	// Auth Routes (Protected)
	auth.Get("/me", middleware.AuthMiddleware(jwtSecret), h.Auth.GetProfile)

	// Product Routes (Protected)
	products := api.Group("/products")
	products.Use(middleware.AuthMiddleware(jwtSecret))
	products.Post("/", h.Product.Create)
	products.Get("/", h.Product.GetAll)
	products.Get("/:id", h.Product.GetByID)
	products.Put("/:id", h.Product.Update)
	products.Delete("/:id", h.Product.Delete)
}
