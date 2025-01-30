package routes

import (
	"health-history/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	personController := controllers.NewPersonController()
	medicationController := controllers.NewMedicationController()
	diagnosticController := controllers.NewDiagnosticController()

	personRoute := r.Group("/api/v1/persons")
	personRoute.GET("/:id", personController.GetPersonById)
	personRoute.GET("/", personController.GetAllPersons)
	personRoute.POST("/", personController.CreatePerson)
	personRoute.PUT("/:id", personController.UpdatePerson)
	personRoute.DELETE("/:id", personController.DeletePerson)

	medicationRoute := r.Group("/api/v1/medications")
	medicationRoute.GET("/:id", medicationController.GetMedicationById)
	medicationRoute.GET("/", medicationController.GetAllMedications)
	medicationRoute.POST("/", medicationController.CreateMedication)
	medicationRoute.PUT("/:id", medicationController.UpdateMedication)
	medicationRoute.DELETE("/:id", medicationController.DeleteMedication)

	diagnosticRoute := r.Group("/api/v1/diagnostics")
	diagnosticRoute.GET("/:id", diagnosticController.GetDiagnosticById)
	diagnosticRoute.GET("/", diagnosticController.GetAllDiagnostics)
	diagnosticRoute.POST("/", diagnosticController.CreateDiagnostic)
	diagnosticRoute.PUT("/:id", diagnosticController.UpdateDiagnostic)
	diagnosticRoute.DELETE("/:id", diagnosticController.DeleteDiagnostic)
}
