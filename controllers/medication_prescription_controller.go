package controllers

import (
	"health-history/models"
	"health-history/repositories"
	"health-history/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicationPrescriptionController struct {
	MedicationPrescriptionRepository *repositories.MedicationPrescriptionRepository
}

func NewMedicationPrescriptionController() *MedicationPrescriptionController {
	return &MedicationPrescriptionController{MedicationPrescriptionRepository: &repositories.MedicationPrescriptionRepository{}}
}

// CreateMedicationPrescription create new medicationprescription
// @Summary Create medicationprescription
// @Description Create a new medicationprescription
// @Tags medicationprescriptions
// @Accept json
// @Produce json
// @Param request body requests.MedicationPrescriptionRequest true "Data of medicationprescription"
// @Success 201 {object} models.MedicationPrescription "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /medicationprescriptions [post]
func (pc *MedicationPrescriptionController) CreateMedicationPrescription(c *gin.Context) {
	var medicationprescriptionRequest requests.MedicationPrescriptionRequest

	if err := c.ShouldBindJSON(&medicationprescriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medicationprescription := models.MedicationPrescription{
		Description: medicationprescriptionRequest.Description,
		PatientID:   medicationprescriptionRequest.PatientID,
		DoctorID:    medicationprescriptionRequest.DoctorID,
	}

	for _, medicationDosage := range medicationprescriptionRequest.MedicationAndDosages {
		medicationprescription.MedicationAndDosages = append(medicationprescription.MedicationAndDosages, models.MedicationDosages{
			MedicationID: medicationDosage.MedicationID,
			Dosage:       medicationDosage.Dosage,
		})
	}

	if err := pc.MedicationPrescriptionRepository.CreateMedicationPrescription(&medicationprescription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medicationprescription)
}

// UpdateMedicationPrescription update a medicationprescription
// @Summary Update  medicationprescription
// @Description Update a exist medicationprescription
// @Tags medicationprescriptions
// @Accept json
// @Produce json
// @Param request body requests.MedicationRequest true "Data of medicationprescription"
// @Param id path int true "MedicationPrescription ID" example(1)
// @Success 200 {object} models.MedicationPrescription "Successful response"
// @Failure 400 {object} map[string]string "Validation error"
// @Router /medicationprescriptions/:id [put]
func (pc *MedicationPrescriptionController) UpdateMedicationPrescription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var medicationprescriptionRequest requests.MedicationPrescriptionRequest
	if err := c.ShouldBindJSON(&medicationprescriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	medicationprescription := models.MedicationPrescription{
		ID:          id,
		Description: medicationprescriptionRequest.Description,
		PatientID:   medicationprescriptionRequest.PatientID,
		DoctorID:    medicationprescriptionRequest.DoctorID,
	}

	for _, medicationDosage := range medicationprescriptionRequest.MedicationAndDosages {
		medicationprescription.MedicationAndDosages = append(medicationprescription.MedicationAndDosages, models.MedicationDosages{
			MedicationID: medicationDosage.MedicationID,
			Dosage:       medicationDosage.Dosage,
		})
	}

	if err := pc.MedicationPrescriptionRepository.UpdateMedicationPrescription(&medicationprescription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicationprescription)
}

// DeleteMedicationPrescription delete a medicationprescription
// @Summary Delete medicationprescription
// @Description Delete a medicationprescription by id
// @Tags medicationprescriptions
// @Produce json
// @Param id path int true "MedicationPrescription ID" example(1)
// @Success 204
// @Router /medicationprescriptions/:id [get]
func (pc *MedicationPrescriptionController) DeleteMedicationPrescription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.MedicationPrescriptionRepository.DeleteMedicationPrescription(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetMedicationPrescriptionById return a medicationprescription
// @Summary Return medicationprescription
// @Description Return medicationprescription by id
// @Tags medicationprescriptions
// @Produce json
// @Param id path int true "MedicationPrescription ID" example(1)
// @Success 200 {object} models.MedicationPrescription "Successful response"
// @Failure 404 {object} map[string]string "Medication prescription not found"
// @Router /medicationprescriptions/:id [get]
func (pc *MedicationPrescriptionController) GetMedicationPrescriptionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	medicationprescription, err := pc.MedicationPrescriptionRepository.GetMedicationPrescriptionById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	c.JSON(http.StatusOK, medicationprescription)
}

// GetAllMedicationPrescriptions return medicationprescription list
// @Summary List medicationprescriptions
// @Description Return medicationprescription list
// @Tags medicationprescriptions
// @Produce json
// @Success 200 {array} models.Medication "Successful response"
// @Router /medicationprescriptions [get]
func (pc *MedicationPrescriptionController) GetAllMedicationPrescriptions(c *gin.Context) {
	medicationprescriptions, err := pc.MedicationPrescriptionRepository.GetAllMedicationPrescriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicationprescriptions)
}
