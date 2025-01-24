package requests

type MedicationRequest struct {
	Name                       string `json:"name" binding:"required"`
	CompanyHoldingRegistration string `json:"CompanyHoldingRegistration" binding:"required"`
	RegulatoryCategory         string `json:"regulatoryCategory" binding:"required"`
	RegistrationNumber         string `json:"registrationNumber" binding:"required"`
	TherapeuticClass           string `json:"therapeuticClass" binding:"required"`
	ActiveIngredient           string `json:"activeIngredient" binding:"required"`
}
