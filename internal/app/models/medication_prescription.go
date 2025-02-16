package models

import "time"

type MedicationPrescription struct {
	ID                   int                 `json:"id"`
	Description          string              `json:"description" binding:"required"`
	PatientID            string              `json:"patient_id" binding:"required"`
	DoctorID             string              `json:"doctor_id" binding:"required"`
	MedicationAndDosages []MedicationDosages `json:"medication_and_dosages" binding:"required" gorm:"type:jsonb"`
	CreatedAt            time.Time           `gorm:"autoCreateTime"`
	UpdatedAt            time.Time           `gorm:"autoUpdateTime"`
}

type MedicationDosages struct {
	MedicationID int    `json:"medication_id" binding:"required"`
	Dosage       string `json:"dosage" binding:"required"`
}
