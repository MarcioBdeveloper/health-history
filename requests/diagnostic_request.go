package requests

type DiagnosticRequest struct {
	Description string   `json:"description" binding:"required" example:"Patient has a mild fever"`
	PatientID   string   `json:"patient_id" binding:"required" example:"1"`
	DoctorID    string   `json:"doctor_id" binding:"required" example:"1"`
	Symptoms    []string `json:"symptoms" binding:"required" example:"['cough', 'headache', 'fever']"`
}
