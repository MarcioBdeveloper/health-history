package main

import (
	"health-history/config"
	"health-history/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	config.ConnectDatabase()
	config.Migrate()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
