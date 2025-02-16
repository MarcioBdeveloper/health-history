package requests

type SymptomRequest struct {
	Name string `json:"name" binding:"required" example:"fever"`
}
