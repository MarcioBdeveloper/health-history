package requests

type MedicationPrescriptionRequest struct {
	Description          string                     `json:"name" binding:"required"`
	PatientID            string                     `json:"patient_id" binding:"required"`
	DoctorID             string                     `json:"doctor_id" binding:"required"`
	MedicationAndDosages []MedicationDosagesRequest `json:"registrationNumber" binding:"required"`
}

type MedicationDosagesRequest struct {
	MedicationID int    `json:"medication_id" binding:"required"`
	Dosage       string `json:"dosage" binding:"required"`
}
