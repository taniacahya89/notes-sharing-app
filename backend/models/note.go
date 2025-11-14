package models

import (
	"time"

	"gorm.io/gorm"
)

// Note represents a note in the system
type Note struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	ImageURL  string         `json:"image_url,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// CreateNoteRequest represents the create note request payload
type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// UpdateNoteRequest represents the update note request payload
type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}