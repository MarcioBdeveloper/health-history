package main

import (
	"health-history/config"
	"health-history/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Failed to load file .env")
		}
	}

	r := gin.Default()
	routes.SetupRoutes(r)
	config.ConnectDatabase()
	config.Migrate()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
