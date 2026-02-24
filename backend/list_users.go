package main

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal("Failed to fetch users:", result.Error)
	}

	fmt.Println("--- Registered Users ---")
	if len(users) == 0 {
		fmt.Println("No users found.")
	}
	for _, user := range users {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Provider: %s\n", user.ID, user.Name, user.Email, user.Provider)
	}
	fmt.Println("------------------------")
}
