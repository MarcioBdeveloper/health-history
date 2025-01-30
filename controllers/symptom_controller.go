package controllers

import (
	"health-history/models"
	"health-history/repositories"
	"health-history/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SymptomController struct {
	SymptomRepository *repositories.SymptomRepository
}

func NewSymptomController() *SymptomController {
	return &SymptomController{SymptomRepository: &repositories.SymptomRepository{}}
}

// CreateSymptom create new symptom
// @Summary Create symptom
// @Description Create a new symptom
// @Tags symptoms
// @Accept json
// @Produce json
// @Param request body requests.SymptomRequest true "Data of symptom"
// @Success 201 {object} models.Symptom
// @Failure 400 {object} map[string]string "Validation error"
// @Router /symptoms [post]
func (pc *SymptomController) CreateSymptom(c *gin.Context) {
	var symptomRequest requests.SymptomRequest

	if err := c.ShouldBindJSON(&symptomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	symptom := models.Symptom{
		Name: symptomRequest.Name,
	}

	if err := pc.SymptomRepository.CreateSymptom(&symptom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, symptom)
}

// UpdateSymptom update a symptom
// @Summary Update  symptom
// @Description Update a exist symptom
// @Tags symptoms
// @Accept json
// @Produce json
// @Param request body requests.SymptomRequest true "Data of symptom"
// @Param id path int true "Symptom ID" example(1)
// @Success 200 {object} models.Symptom
// @Router /symptoms/:id [put]
func (pc *SymptomController) UpdateSymptom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var symptomRequest requests.SymptomRequest
	if err := c.ShouldBindJSON(&symptomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	symptom := models.Symptom{
		ID:   id,
		Name: symptomRequest.Name,
	}

	if err := pc.SymptomRepository.UpdateSymptom(&symptom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, symptom)
}

// DeleteSymptom delete a symptom
// @Summary Delete symptom
// @Description Delete a symptom by id
// @Tags symptoms
// @Produce json
// @Param id path int true "Symptom ID" example(1)
// @Success 204
// @Router /symptoms/:id [get]
func (pc *SymptomController) DeleteSymptom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.SymptomRepository.DeleteSymptom(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetSymptomById return a symptom
// @Summary Return symptom
// @Description Return symptom by id
// @Tags symptoms
// @Produce json
// @Param id path int true "Symptom ID" example(1)
// @Success 200 {object} models.Symptom
// @Failure 404 {object} map[string]string "Symptom not found"
// @Router /symptoms/:id [get]
func (pc *SymptomController) GetSymptomById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	symptom, err := pc.SymptomRepository.GetSymptomById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Symptom not found"})
		return
	}

	c.JSON(http.StatusOK, symptom)
}

// GetAllSymptoms return symptom list
// @Summary List symptoms
// @Description Return person list
// @Tags symptoms
// @Produce json
// @Success 200 {array} models.Symptom
// @Router /symptoms [get]
func (pc *SymptomController) GetAllSymptoms(c *gin.Context) {
	symptoms, err := pc.SymptomRepository.GetAllSymptoms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, symptoms)
}
