package repositories

import (
	"health-history/internal/app/config"
	"health-history/internal/app/models"
)

type SymptomRepository struct{}

func (r *SymptomRepository) CreateSymptom(symptom *models.Symptom) error {
	result := config.DB.Create(symptom)
	return result.Error
}

func (r *SymptomRepository) UpdateSymptom(symptom *models.Symptom) error {
	result := config.DB.Save(symptom)
	return result.Error
}

func (r *SymptomRepository) DeleteSymptom(id int) error {
	result := config.DB.Delete(&models.Symptom{}, id)
	return result.Error
}

func (r *SymptomRepository) GetSymptomById(id int) (*models.Symptom, error) {
	var symptom models.Symptom
	result := config.DB.First(&symptom, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &symptom, nil
}

func (r *SymptomRepository) GetAllSymptoms() ([]models.Symptom, error) {
	var symptoms []models.Symptom
	result := config.DB.Find(&symptoms)
	return symptoms, result.Error
}
