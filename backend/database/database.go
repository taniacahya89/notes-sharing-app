package database

import (
	"log"
	"notes-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB connects to the PostgreSQL database
func ConnectDB(databaseURL string) {
	var err error

	// Connect to database
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")

	// Seed dummy data
	seedData()
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error getting database instance:", err)
		return
	}
	sqlDB.Close()
}

// seedData inserts dummy data for testing
func seedData() {
	// Check if data already exists
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Database already contains data, skipping seed")
		return
	}

	// Create dummy users
	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "$2a$10$5JZKqx5JqZKqx5JqZKqx5OQKqx5JqZKqx5JqZKqx5JqZKqx5JqZK", // "password123"
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: "$2a$10$5JZKqx5JqZKqx5JqZKqx5OQKqx5JqZKqx5JqZKqx5JqZKqx5JqZK", // "password123"
		},
	}

	for _, user := range users {
		DB.Create(&user)
	}

	// Create dummy notes
	notes := []models.Note{
		{
			UserID:   1,
			Title:    "Welcome Note",
			Content:  "This is your first note! Welcome to Notes Sharing App.",
			ImageURL: "",
		},
		{
			UserID:   1,
			Title:    "Shopping List",
			Content:  "1. Milk\n2. Bread\n3. Eggs\n4. Butter",
			ImageURL: "",
		},
		{
			UserID:   2,
			Title:    "Meeting Notes",
			Content:  "Discussed project timeline and deliverables for Q4 2024.",
			ImageURL: "",
		},
	}

	for _, note := range notes {
		DB.Create(&note)
	}

	log.Println("Dummy data seeded successfully")
}