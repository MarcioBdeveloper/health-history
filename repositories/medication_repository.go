package repositories

import (
	"health-history/config"
	"health-history/models"
)

type MedicationRepository struct{}

func (r *MedicationRepository) CreateMedication(medication *models.Medication) error {
	result := config.DB.Create(medication)
	return result.Error
}

func (r *MedicationRepository) UpdateMedication(medication *models.Medication) error {
	result := config.DB.Save(medication)
	return result.Error
}

func (r *MedicationRepository) DeleteMedication(id int) error {
	result := config.DB.Delete(&models.Medication{}, id)
	return result.Error
}

func (r *MedicationRepository) GetMedicationById(id int) (*models.Medication, error) {
	var medication models.Medication
	result := config.DB.First(&medication, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &medication, nil
}

func (r *MedicationRepository) GetAllMedications() ([]models.Medication, error) {
	var medications []models.Medication
	result := config.DB.Find(&medications)
	return medications, result.Error
}
