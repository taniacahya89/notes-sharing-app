package main

import (
	"fmt"
	"log"
	"notes-app/config"
	"notes-app/database"
	"notes-app/routes"
	"notes-app/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	utils.InitLogger()
	utils.LogInfo("Starting Notes Sharing App...")

	// Initialize database
	database.ConnectDB(cfg.DatabaseURL)
	defer database.CloseDB()

	// Create directories for uploads and logs
	createDirectories()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Serve static files (uploads)
	app.Static("/uploads", "./uploads")

	// Setup routes
	routes.SetupRoutes(app)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Notes API is running",
		})
	})

	// Start server
	port := cfg.Port
	utils.LogInfo(fmt.Sprintf("Server is running on port %s", port))
	log.Fatal(app.Listen(":" + port))
}

// customErrorHandler handles errors globally
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	utils.LogError(fmt.Sprintf("Error: %v", err))

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// createDirectories creates necessary directories
func createDirectories() {
	dirs := []string{"./uploads", "./logs"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
	}
}