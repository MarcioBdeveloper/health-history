package models

import "time"

type Diagnostic struct {
	ID          int       `json:"id"`
	Description string    `json:"description" binding:"required"`
	PatientID   string    `json:"patient_id" binding:"required"`
	DoctorID    string    `json:"doctor_id" binding:"required"`
	Symptoms    []string  `json:"symptoms" binding:"required" gorm:"type:jsonb"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
