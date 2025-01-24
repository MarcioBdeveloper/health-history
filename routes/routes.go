package routes

import (
	"health-history/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	personController := controllers.NewPersonController()

	personRoute := r.Group("/persons")
	personRoute.GET("/:id", personController.GetPersonById)
	personRoute.GET("/", personController.GetAllPersons)
	personRoute.POST("/", personController.CreatePerson)
	personRoute.PUT("/:id", personController.UpdatePerson)
	personRoute.DELETE("/:id", personController.DeletePerson)
}
