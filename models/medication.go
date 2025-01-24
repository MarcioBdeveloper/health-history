package models

import "time"

type Medication struct {
	ID                         int       `json:"id"`
	Name                       string    `json:"name" binding:"required"`
	CompanyHoldingRegistration string    `json:"companyHoldingRegistration" binding:"required"`
	RegulatoryCategory         string    `json:"regulatoryCategory" binding:"required"`
	RegistrationNumber         string    `json:"registrationNumber" binding:"required"`
	TherapeuticClass           string    `json:"therapeuticClass" binding:"required"`
	ActiveIngredient           string    `json:"activeIngredient" binding:"required"`
	CreatedAt                  time.Time `gorm:"autoCreateTime"`
	UpdatedAt                  time.Time `gorm:"autoUpdateTime"`
}
