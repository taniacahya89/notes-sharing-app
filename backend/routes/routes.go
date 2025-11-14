package routes

import (
	"notes-app/handlers"
	"notes-app/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all application routes
func SetupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api")

	// Auth routes (no authentication required)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Notes routes (authentication required)
	notes := api.Group("/notes", middleware.AuthMiddleware)
	notes.Get("/", handlers.GetNotes)
	notes.Post("/", handlers.CreateNote)
	notes.Get("/:id", handlers.GetNote)
	notes.Put("/:id", handlers.UpdateNote)
	notes.Delete("/:id", handlers.DeleteNote)
	notes.Post("/:id/upload", handlers.UploadImage)
}