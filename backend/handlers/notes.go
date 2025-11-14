package handlers

import (
	"fmt"
	"notes-app/database"
	"notes-app/models"
	"notes-app/utils"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetNotes retrieves all notes for the authenticated user
func GetNotes(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var notes []models.Note
	if err := database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notes).Error; err != nil {
		utils.LogError("Failed to get notes: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve notes",
		})
	}

	return c.JSON(fiber.Map{
		"notes": notes,
	})
}

// GetNote retrieves a specific note by ID
func GetNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.JSON(note)
}

// CreateNote creates a new note
func CreateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req models.CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		utils.LogError("Failed to parse create note request: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Title == "" || req.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and content are required",
		})
	}

	// Create note
	note := models.Note{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := database.DB.Create(&note).Error; err != nil {
		utils.LogError("Failed to create note: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create note",
		})
	}

	utils.LogInfo(fmt.Sprintf("Note created: ID=%d, UserID=%d", note.ID, userID))

	return c.Status(fiber.StatusCreated).JSON(note)
}

// UpdateNote updates an existing note
func UpdateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	var req models.UpdateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		utils.LogError("Failed to parse update note request: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update fields if provided
	if req.Title != "" {
		note.Title = req.Title
	}
	if req.Content != "" {
		note.Content = req.Content
	}

	if err := database.DB.Save(&note).Error; err != nil {
		utils.LogError("Failed to update note: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	utils.LogInfo(fmt.Sprintf("Note updated: ID=%d, UserID=%d", note.ID, userID))

	return c.JSON(note)
}

// DeleteNote deletes a note
func DeleteNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	if err := database.DB.Delete(&note).Error; err != nil {
		utils.LogError("Failed to delete note: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete note",
		})
	}

	utils.LogInfo(fmt.Sprintf("Note deleted: ID=%d, UserID=%d", note.ID, userID))

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}

// UploadImage uploads an image for a note
func UploadImage(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	// Check if note exists and belongs to user
	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	// Get uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No image file provided",
		})
	}

	// Validate file type
	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !allowedExts[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid file type. Only jpg, jpeg, png, and gif are allowed",
		})
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%d%s", time.Now().Unix(), note.ID, ext)
	filepath := fmt.Sprintf("./uploads/%s", filename)

	// Save file
	if err := c.SaveFile(file, filepath); err != nil {
		utils.LogError("Failed to save image: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save image",
		})
	}

	// Update note with image URL
	note.ImageURL = "/uploads/" + filename
	if err := database.DB.Save(&note).Error; err != nil {
		utils.LogError("Failed to update note with image URL: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	utils.LogInfo(fmt.Sprintf("Image uploaded for note: ID=%d, UserID=%d", note.ID, userID))

	return c.JSON(note)
}
