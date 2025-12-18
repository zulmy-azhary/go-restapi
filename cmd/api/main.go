package main

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/database"
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/routes"
	"go-rest-api/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Initialize repositories
	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWT)
	productService := services.NewProductService(productRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	healthHandler := handlers.NewHealthHandler(db)
	
	// Group handlers
	handlers := handlers.NewHandlers(authHandler, productHandler, healthHandler)

	// Setup routes
	routes.Setup(app, handlers, cfg.JWT.Secret)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
