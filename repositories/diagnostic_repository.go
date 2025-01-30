package repositories

import (
	"health-history/config"
	"health-history/models"
)

type DiagnosticRepository struct{}

func (r *DiagnosticRepository) CreateDiagnostic(diagnostic *models.Diagnostic) error {
	result := config.DB.Create(diagnostic)
	return result.Error
}

func (r *DiagnosticRepository) UpdateDiagnostic(diagnostic *models.Diagnostic) error {
	result := config.DB.Save(diagnostic)
	return result.Error
}

func (r *DiagnosticRepository) DeleteDiagnostic(id int) error {
	result := config.DB.Delete(&models.Diagnostic{}, id)
	return result.Error
}

func (r *DiagnosticRepository) GetDiagnosticById(id int) (*models.Diagnostic, error) {
	var diagnostic models.Diagnostic
	result := config.DB.First(&diagnostic, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &diagnostic, nil
}

func (r *DiagnosticRepository) GetAllDiagnostics() ([]models.Diagnostic, error) {
	var medications []models.Diagnostic
	result := config.DB.Find(&medications)
	return medications, result.Error
}
