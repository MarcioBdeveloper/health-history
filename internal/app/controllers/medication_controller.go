package controllers

import (
	"health-history/internal/app/models"
	"health-history/internal/app/repositories"
	"health-history/internal/app/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicationController struct {
	MedicationRepository *repositories.MedicationRepository
}

func NewMedicationController() *MedicationController {
	return &MedicationController{MedicationRepository: &repositories.MedicationRepository{}}
}

// CreateMedication create new medication
// @Summary Create medication
// @Description Create a new medication
// @Tags medications
// @Accept json
// @Produce json
// @Param request body requests.MedicationRequest true "Data of medication"
// @Success 201 {object} models.Medication "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /medications [post]
func (pc *MedicationController) CreateMedication(c *gin.Context) {
	var medicationRequest requests.MedicationRequest

	if err := c.ShouldBindJSON(&medicationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medication := models.Medication{
		Name:                       medicationRequest.Name,
		CompanyHoldingRegistration: medicationRequest.CompanyHoldingRegistration,
		RegulatoryCategory:         medicationRequest.RegulatoryCategory,
		RegistrationNumber:         medicationRequest.RegistrationNumber,
		TherapeuticClass:           medicationRequest.TherapeuticClass,
		ActiveIngredient:           medicationRequest.ActiveIngredient,
	}

	if err := pc.MedicationRepository.CreateMedication(&medication); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medication)
}

// UpdateMedication update a medication
// @Summary Update  medication
// @Description Update a exist medication
// @Tags medications
// @Accept json
// @Produce json
// @Param request body requests.MedicationRequest true "Data of medication"
// @Param id path int true "Medication ID" example(1)
// @Success 200 {object} models.Medication "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /medications/:id [put]
func (pc *MedicationController) UpdateMedication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var medicationRequest requests.MedicationRequest
	if err := c.ShouldBindJSON(&medicationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medication := models.Medication{
		ID:                         id,
		Name:                       medicationRequest.Name,
		CompanyHoldingRegistration: medicationRequest.CompanyHoldingRegistration,
		RegulatoryCategory:         medicationRequest.RegulatoryCategory,
		RegistrationNumber:         medicationRequest.RegistrationNumber,
		TherapeuticClass:           medicationRequest.TherapeuticClass,
		ActiveIngredient:           medicationRequest.ActiveIngredient,
	}

	if err := pc.MedicationRepository.UpdateMedication(&medication); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)
}

// DeleteMedication delete a medication
// @Summary Delete medication
// @Description Delete a medication by id
// @Tags medications
// @Produce json
// @Param id path int true "Medication ID" example(1)
// @Success 204
// @Router /medications/:id [get]
func (pc *MedicationController) DeleteMedication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.MedicationRepository.DeleteMedication(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMedicationById return a medication
// @Summary Return medication
// @Description Return medication by id
// @Tags medications
// @Produce json
// @Param id path int true "Medication ID" example(1)
// @Success 200 {object} models.Medication "Successful response"
// @Failure 404 {object} map[string]string "Medication not found"
// @Router /medications/:id [get]
func (pc *MedicationController) GetMedicationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	medication, err := pc.MedicationRepository.GetMedicationById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	c.JSON(http.StatusOK, medication)
}

// GetAllMedications return medication list
// @Summary List medications
// @Description Return medication list
// @Tags medications
// @Produce json
// @Success 200 {array} models.Medication "Successful response"
// @Router /medications [get]
func (pc *MedicationController) GetAllMedications(c *gin.Context) {
	medications, err := pc.MedicationRepository.GetAllMedications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medications)
}
