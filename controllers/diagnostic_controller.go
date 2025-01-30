package controllers

import (
	"health-history/models"
	"health-history/repositories"
	"health-history/requests"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiagnosticController struct {
	DiagnosticRepository *repositories.DiagnosticRepository
}

func NewDiagnosticController() *DiagnosticController {
	return &DiagnosticController{DiagnosticRepository: &repositories.DiagnosticRepository{}}
}

// CreateDiagnostic create new diagnostic
// @Summary Create diagnostic
// @Description Create a new diagnostic
// @Tags diagnostics
// @Accept json
// @Produce json
// @Param request body requests.DiagnosticRequest true "Data of diagnostic"
// @Success 201 {object} models.Diagnostic
// @Failure 400 {object} map[string]string "Validation error"
// @Router /diagnostics [post]
func (pc *DiagnosticController) CreateDiagnostic(c *gin.Context) {
	var diagnosticRequest requests.DiagnosticRequest

	if err := c.ShouldBindJSON(&diagnosticRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	diagnostic := models.Diagnostic{
		Description: diagnosticRequest.Description,
		PatientID:   diagnosticRequest.PatientID,
		DoctorID:    diagnosticRequest.DoctorID,
		Symptoms:    diagnosticRequest.Symptoms,
	}

	if err := pc.DiagnosticRepository.CreateDiagnostic(&diagnostic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, diagnostic)
}

// UpdateDiagnostic update a diagnostic
// @Summary Update  diagnostic
// @Description Update a exist diagnostic
// @Tags diagnostics
// @Accept json
// @Produce json
// @Param request body requests.DiagnosticRequest true "Data of diagnostic"
// @Param id path int true "Diagnostic ID" example(1)
// @Success 200 {object} models.Diagnostic
// @Router /diagnostics/:id [put]
func (pc *DiagnosticController) UpdateDiagnostic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var diagnosticRequest requests.DiagnosticRequest
	if err := c.ShouldBindJSON(&diagnosticRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	diagnostic := models.Diagnostic{
		ID:          id,
		Description: diagnosticRequest.Description,
		PatientID:   diagnosticRequest.PatientID,
		DoctorID:    diagnosticRequest.DoctorID,
		Symptoms:    diagnosticRequest.Symptoms,
	}

	if err := pc.DiagnosticRepository.UpdateDiagnostic(&diagnostic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnostic)
}

// DeleteDiagnostic delete a diagnostic
// @Summary Delete diagnostic
// @Description Delete a diagnostic by id
// @Tags diagnostics
// @Produce json
// @Param id path int true "Diagnostic ID" example(1)
// @Success 204
// @Router /diagnostics/:id [get]
func (pc *DiagnosticController) DeleteDiagnostic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.DiagnosticRepository.DeleteDiagnostic(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetDiagnosticById return a diagnostic
// @Summary Return diagnostic
// @Description Return diagnostic by id
// @Tags diagnostics
// @Produce json
// @Param id path int true "Diagnostic ID" example(1)
// @Success 200 {object} models.Diagnostic
// @Failure 404 {object} map[string]string "Diagnostic not found"
// @Router /diagnostics/:id [get]
func (pc *DiagnosticController) GetDiagnosticById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	diagnostic, err := pc.DiagnosticRepository.GetDiagnosticById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diagnostic not found"})
		return
	}

	c.JSON(http.StatusOK, diagnostic)
}

// GetAllDiagnostics return diagnostic list
// @Summary List diagnostics
// @Description Return person list
// @Tags diagnostics
// @Produce json
// @Success 200 {array} models.Diagnostic
// @Router /diagnostics [get]
func (pc *DiagnosticController) GetAllDiagnostics(c *gin.Context) {
	diagnostics, err := pc.DiagnosticRepository.GetAllDiagnostics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnostics)
}
