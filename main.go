package main

import (
	"health-history/config"
	"health-history/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	config.ConnectDatabase()
	config.Migrate()

	r.Run(":8080")
}
