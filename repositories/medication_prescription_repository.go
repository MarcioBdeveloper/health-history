package repositories

import (
	"health-history/config"
	"health-history/models"
)

type MedicationPrescriptionRepository struct{}

func (r *MedicationPrescriptionRepository) CreateMedicationPrescription(medicationPrescription *models.MedicationPrescription) error {
	result := config.DB.Create(medicationPrescription)
	return result.Error
}

func (r *MedicationPrescriptionRepository) UpdateMedicationPrescription(medicationPrescription *models.MedicationPrescription) error {
	result := config.DB.Save(medicationPrescription)
	return result.Error
}

func (r *MedicationPrescriptionRepository) DeleteMedicationPrescription(id int) error {
	result := config.DB.Delete(&models.MedicationPrescription{}, id)
	return result.Error
}

func (r *MedicationPrescriptionRepository) GetMedicationPrescriptionById(id int) (*models.MedicationPrescription, error) {
	var medicationPrescription models.MedicationPrescription
	result := config.DB.First(&medicationPrescription, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &medicationPrescription, nil
}

func (r *MedicationPrescriptionRepository) GetAllMedicationPrescriptions() ([]models.MedicationPrescription, error) {
	var medicationPrescriptions []models.MedicationPrescription
	result := config.DB.Find(&medicationPrescriptions)
	return medicationPrescriptions, result.Error
}
