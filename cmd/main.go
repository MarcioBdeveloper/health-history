package main

import (
	_ "health-history/docs"
	"health-history/internal/app/config"
	"health-history/internal/app/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Health History API
// @version         1.0
// @description     The objective this application is save history medicaments of person.

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.SetupRoutes(r)
	config.ConnectDatabase()
	config.Migrate()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
