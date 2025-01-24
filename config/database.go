package config

import (
	"fmt"
	"health-history/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsnDefault := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	databaseName := "db_health_history"

	tempDB, err := gorm.Open(postgres.Open(dsnDefault), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := tempDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", databaseName)).Error; err != nil {
		fmt.Println("Database already exists or cannot be created:", err)
	}

	errload := godotenv.Load()
	if errload != nil {
		log.Printf("Warning: Could not load .env file. Using system environment variables.")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations executed successfully!")
}
